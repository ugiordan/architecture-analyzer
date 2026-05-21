package extractor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractPythonExternalConnections_Database(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app/db.py", `
import psycopg2

def get_connection():
    conn = psycopg2.connect("postgresql://user:pass@localhost:5432/mydb")
    return conn

def get_redis():
    r = redis.Redis(host="localhost", port=6379)
    return r
`)

	conns := extractPythonExternalConnections(dir)

	byService := map[string][]ExternalConnection{}
	for _, c := range conns {
		byService[c.Service] = append(byService[c.Service], c)
	}

	if len(byService["postgres"]) < 1 {
		t.Error("expected at least 1 postgres connection")
	}
	if len(byService["redis"]) < 1 {
		t.Error("expected at least 1 redis connection")
	}

	for _, c := range byService["postgres"] {
		if c.Type != "database" {
			t.Errorf("postgres connection type = %q, want database", c.Type)
		}
		if c.Target != "" && contains(c.Target, "pass") {
			t.Errorf("target should redact password, got %q", c.Target)
		}
	}
}

func TestExtractPythonExternalConnections_ObjectStorage(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "storage.py", `
import boto3

def upload_model():
    s3 = boto3.client('s3')
    s3.upload_file("model.bin", "my-bucket", "models/model.bin")
`)

	conns := extractPythonExternalConnections(dir)

	found := false
	for _, c := range conns {
		if c.Service == "s3" && c.Type == "object-storage" {
			found = true
		}
	}
	if !found {
		t.Error("expected s3 object-storage connection")
	}
}

func TestExtractPythonExternalConnections_GRPC(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "client.py", `
import grpc

def connect():
    channel = grpc.insecure_channel("localhost:50051")
    return channel
`)

	conns := extractPythonExternalConnections(dir)

	found := false
	for _, c := range conns {
		if c.Service == "grpc" && c.Type == "grpc" && c.Target == "localhost:50051" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected grpc connection with target localhost:50051, got %v", conns)
	}
}

func TestExtractPythonExternalConnections_Messaging(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "producer.py", `
from kafka import KafkaProducer

def send_event():
    producer = KafkaProducer(bootstrap_servers="kafka:9092")
    producer.send("events", b"hello")
`)

	conns := extractPythonExternalConnections(dir)

	found := false
	for _, c := range conns {
		if c.Service == "kafka" && c.Type == "messaging" {
			found = true
		}
	}
	if !found {
		t.Error("expected kafka messaging connection")
	}
}

func TestExtractPythonExternalConnections_HTTPClients(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "api.py", `
import requests
import httpx

def fetch_data():
    resp = requests.get("https://api.example.com/data")
    return resp.json()

async def fetch_async():
    async with httpx.AsyncClient() as client:
        resp = await client.get("https://api.example.com/data")
`)

	conns := extractPythonExternalConnections(dir)

	services := map[string]bool{}
	for _, c := range conns {
		services[c.Service] = true
	}
	if !services["requests"] {
		t.Error("expected requests HTTP client connection")
	}
	if !services["httpx"] {
		t.Error("expected httpx HTTP client connection")
	}
}

func TestExtractPythonExternalConnections_SDKClients(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "llm.py", `
from openai import OpenAI
from llama_stack_client import LlamaStackClient
import chromadb

def get_openai():
    client = OpenAI()
    return client

def get_vector_store():
    ls = LlamaStackClient()
    chroma = chromadb.HttpClient()
    return ls, chroma
`)

	conns := extractPythonExternalConnections(dir)

	services := map[string]bool{}
	for _, c := range conns {
		services[c.Service] = true
	}
	for _, want := range []string{"openai", "llama-stack", "chromadb"} {
		if !services[want] {
			t.Errorf("expected %s SDK connection, got services: %v", want, services)
		}
	}
}

func TestExtractPythonExternalConnections_SkipsTests(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "test_db.py", `
import psycopg2
conn = psycopg2.connect("postgresql://localhost/test")
`)
	writeFile(t, dir, "tests/conftest.py", `
import redis
r = redis.Redis()
`)

	conns := extractPythonExternalConnections(dir)
	if len(conns) != 0 {
		t.Errorf("expected 0 connections from test files, got %d: %v", len(conns), conns)
	}
}

func TestExtractPythonExternalConnections_SQLAlchemy(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "models.py", `
from sqlalchemy import create_engine

engine = create_engine("postgresql://user:secret@db.example.com:5432/app")
`)

	conns := extractPythonExternalConnections(dir)

	found := false
	for _, c := range conns {
		if c.Service == "postgres" && c.Type == "database" {
			found = true
			if contains(c.Target, "secret") {
				t.Errorf("target should redact password, got %q", c.Target)
			}
		}
	}
	if !found {
		t.Error("expected postgres connection from create_engine")
	}
}

func TestExtractPythonExternalConnections_FunctionContext(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "service.py", `
class MyService:
    def connect(self):
        self.db = psycopg2.connect("postgresql://localhost/db")

def standalone():
    requests.get("https://api.example.com")
`)

	conns := extractPythonExternalConnections(dir)

	funcNames := map[string]bool{}
	for _, c := range conns {
		if c.Function != "" {
			funcNames[c.Function] = true
		}
	}
	if !funcNames["connect"] && !funcNames["MyService"] {
		t.Errorf("expected function context for db connection, got functions: %v", funcNames)
	}
}

func writeFile(t *testing.T, dir, rel, content string) {
	t.Helper()
	abs := filepath.Join(dir, rel)
	os.MkdirAll(filepath.Dir(abs), 0o755)
	if err := os.WriteFile(abs, []byte(content), 0o644); err != nil {
		t.Fatalf("writing %s: %v", rel, err)
	}
}

func contains(s, sub string) bool {
	return len(s) > 0 && len(sub) > 0 && len(s) >= len(sub) && indexOf(s, sub) >= 0
}

func indexOf(s, sub string) int {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
