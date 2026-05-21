package extractor

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Python database connection patterns.
var pyDBPatterns = []struct {
	re       *regexp.Regexp
	service  string
	connType string
}{
	// PostgreSQL
	{regexp.MustCompile(`["'](postgres(?:ql)?://[^"']+)["']`), "postgres", "database"},
	{regexp.MustCompile(`psycopg2?\.connect\(`), "postgres", "database"},
	{regexp.MustCompile(`asyncpg\.connect\(`), "postgres", "database"},
	{regexp.MustCompile(`create_engine\(\s*["']postgres`), "postgres", "database"},
	{regexp.MustCompile(`create_async_engine\(\s*["']postgres`), "postgres", "database"},

	// MySQL
	{regexp.MustCompile(`["'](mysql(?:\+pymysql)?://[^"']+)["']`), "mysql", "database"},
	{regexp.MustCompile(`pymysql\.connect\(`), "mysql", "database"},
	{regexp.MustCompile(`MySQLdb\.connect\(`), "mysql", "database"},
	{regexp.MustCompile(`create_engine\(\s*["']mysql`), "mysql", "database"},

	// MongoDB
	{regexp.MustCompile(`["'](mongodb(?:\+srv)?://[^"']+)["']`), "mongodb", "database"},
	{regexp.MustCompile(`(?:pymongo\.)?MongoClient\(`), "mongodb", "database"},
	{regexp.MustCompile(`AsyncIOMotorClient\(`), "mongodb", "database"},

	// Redis
	{regexp.MustCompile(`["'](redis(?:s)?://[^"']+)["']`), "redis", "database"},
	{regexp.MustCompile(`redis\.Redis\(`), "redis", "database"},
	{regexp.MustCompile(`redis\.StrictRedis\(`), "redis", "database"},
	{regexp.MustCompile(`aioredis\.from_url\(`), "redis", "database"},
	{regexp.MustCompile(`redis\.asyncio\.Redis\(`), "redis", "database"},

	// SQLite
	{regexp.MustCompile(`sqlite3\.connect\(`), "sqlite", "database"},
	{regexp.MustCompile(`create_engine\(\s*["']sqlite`), "sqlite", "database"},

	// SQLAlchemy generic
	{regexp.MustCompile(`create_engine\(`), "sqlalchemy", "database"},
	{regexp.MustCompile(`create_async_engine\(`), "sqlalchemy", "database"},
}

// Python object storage patterns.
var pyStoragePatterns = []struct {
	re      *regexp.Regexp
	service string
}{
	// AWS S3 via boto3
	{regexp.MustCompile(`boto3\.client\(\s*["']s3["']`), "s3"},
	{regexp.MustCompile(`boto3\.resource\(\s*["']s3["']`), "s3"},
	{regexp.MustCompile(`["'](s3://[^"']+)["']`), "s3"},
	{regexp.MustCompile(`s3fs\.S3FileSystem\(`), "s3"},

	// MinIO
	{regexp.MustCompile(`(?:Minio|MinIO)\(`), "minio"},

	// GCS
	{regexp.MustCompile(`storage\.Client\(`), "gcs"},
	{regexp.MustCompile(`["'](gs://[^"']+)["']`), "gcs"},

	// Azure Blob
	{regexp.MustCompile(`BlobServiceClient\(`), "azure-blob"},
	{regexp.MustCompile(`blob\.core\.windows\.net`), "azure-blob"},
}

// Python gRPC patterns.
var pyGRPCPatterns = []*regexp.Regexp{
	regexp.MustCompile(`grpc\.insecure_channel\(\s*["']([^"']+)["']`),
	regexp.MustCompile(`grpc\.secure_channel\(\s*["']([^"']+)["']`),
	regexp.MustCompile(`grpc\.aio\.insecure_channel\(\s*["']([^"']+)["']`),
	regexp.MustCompile(`grpc\.aio\.secure_channel\(\s*["']([^"']+)["']`),
}

// Python messaging patterns.
var pyMsgPatterns = []struct {
	re      *regexp.Regexp
	service string
}{
	// Kafka
	{regexp.MustCompile(`KafkaProducer\(`), "kafka"},
	{regexp.MustCompile(`KafkaConsumer\(`), "kafka"},
	{regexp.MustCompile(`confluent_kafka\.Producer\(`), "kafka"},
	{regexp.MustCompile(`confluent_kafka\.Consumer\(`), "kafka"},

	// NATS
	{regexp.MustCompile(`nats\.connect\(\s*["']([^"']+)["']`), "nats"},

	// RabbitMQ / AMQP
	{regexp.MustCompile(`["'](amqp(?:s)?://[^"']+)["']`), "rabbitmq"},
	{regexp.MustCompile(`pika\.BlockingConnection\(`), "rabbitmq"},
	{regexp.MustCompile(`pika\.SelectConnection\(`), "rabbitmq"},
}

// Python HTTP client patterns.
var pyHTTPPatterns = []struct {
	re      *regexp.Regexp
	service string
}{
	{regexp.MustCompile(`requests\.(get|post|put|delete|patch|head|options)\(`), "requests"},
	{regexp.MustCompile(`requests\.Session\(`), "requests"},
	{regexp.MustCompile(`httpx\.Client\(`), "httpx"},
	{regexp.MustCompile(`httpx\.AsyncClient\(`), "httpx"},
	{regexp.MustCompile(`httpx\.(get|post|put|delete|patch)\(`), "httpx"},
	{regexp.MustCompile(`urllib\.request\.urlopen\(`), "urllib"},
	{regexp.MustCompile(`aiohttp\.ClientSession\(`), "aiohttp"},
}

// Python LLM/ML SDK client patterns.
var pySDKPatterns = []struct {
	re      *regexp.Regexp
	service string
}{
	// LLM providers
	{regexp.MustCompile(`(?:openai\.)?OpenAI\(`), "openai"},
	{regexp.MustCompile(`(?:openai\.)?AsyncOpenAI\(`), "openai"},
	{regexp.MustCompile(`(?:anthropic\.)?Anthropic\(`), "anthropic"},
	{regexp.MustCompile(`(?:anthropic\.)?AsyncAnthropic\(`), "anthropic"},
	{regexp.MustCompile(`LlamaStackClient\(`), "llama-stack"},
	{regexp.MustCompile(`(?:cohere\.)?Client\(\s*api_key`), "cohere"},

	// Vector databases
	{regexp.MustCompile(`chromadb\.Client\(`), "chromadb"},
	{regexp.MustCompile(`chromadb\.HttpClient\(`), "chromadb"},
	{regexp.MustCompile(`chromadb\.PersistentClient\(`), "chromadb"},
	{regexp.MustCompile(`(?:elasticsearch\.)?Elasticsearch\(`), "elasticsearch"},
	{regexp.MustCompile(`(?:weaviate\.)?Client\(\s*url`), "weaviate"},
	{regexp.MustCompile(`(?:pinecone\.)?Pinecone\(`), "pinecone"},
	{regexp.MustCompile(`QdrantClient\(`), "qdrant"},
	{regexp.MustCompile(`MilvusClient\(`), "milvus"},

	// ML platforms
	{regexp.MustCompile(`mlflow\.(set_tracking_uri|start_run|log_)`), "mlflow"},
	{regexp.MustCompile(`wandb\.(init|log)\(`), "wandb"},
}

var pyCommentRE = regexp.MustCompile(`#.*$`)
var pyStringLitRE = regexp.MustCompile(`(?:"(?:[^"\\]|\\.)*"|'(?:[^'\\]|\\.)*')`)
var pyFuncDefRE = regexp.MustCompile(`^\s*(?:async\s+)?def\s+(\w+)\s*\(`)
var pyClassDefRE = regexp.MustCompile(`^\s*class\s+(\w+)`)

func extractPythonExternalConnections(repoPath string) []ExternalConnection {
	var connections []ExternalConnection

	pyFiles := findFiles(repoPath, []string{"**/*.py"})

	for _, fpath := range pyFiles {
		base := strings.ToLower(fpath)
		if strings.Contains(base, "test_") || strings.HasSuffix(base, "_test.py") ||
			strings.Contains(fpath, "/tests/") || strings.Contains(fpath, "/test/") {
			continue
		}

		info, err := os.Lstat(fpath)
		if err != nil || info.Mode()&os.ModeSymlink != 0 {
			continue
		}
		if info.Size() > maxFileSize {
			log.Printf("skipping oversized file %s: %d bytes", fpath, info.Size())
			continue
		}
		data, err := os.ReadFile(fpath)
		if err != nil {
			continue
		}

		content := string(data)
		lines := strings.Split(content, "\n")
		source := relativePath(repoPath, fpath)

		currentFunc := ""
		indentStack := []int{}

		for lineNum, line := range lines {
			stripped := strings.TrimSpace(line)
			if stripped == "" || strings.HasPrefix(stripped, "#") {
				continue
			}

			// Track function context via indentation
			if fn := pyFuncDefRE.FindStringSubmatch(line); fn != nil {
				currentFunc = fn[1]
				indentStack = append(indentStack, leadingSpaces(line))
			} else if cls := pyClassDefRE.FindStringSubmatch(line); cls != nil {
				currentFunc = cls[1]
				indentStack = append(indentStack, leadingSpaces(line))
			} else if len(indentStack) > 0 && len(stripped) > 0 {
				indent := leadingSpaces(line)
				for len(indentStack) > 0 && indent <= indentStack[len(indentStack)-1] {
					indentStack = indentStack[:len(indentStack)-1]
					currentFunc = ""
				}
			}

			loc := source + ":" + strconv.Itoa(lineNum+1)

			// Database
			for _, p := range pyDBPatterns {
				matches := p.re.FindStringSubmatch(stripped)
				if matches == nil {
					continue
				}
				target := ""
				if len(matches) > 1 && strings.Contains(matches[1], "://") {
					target = redactConnectionString(matches[1])
				}
				connections = append(connections, ExternalConnection{
					Type:     p.connType,
					Service:  p.service,
					Target:   target,
					Source:   loc,
					Function: currentFunc,
				})
			}

			// Object storage
			for _, p := range pyStoragePatterns {
				matches := p.re.FindStringSubmatch(stripped)
				if matches == nil {
					continue
				}
				target := ""
				if len(matches) > 1 && strings.Contains(matches[1], "://") {
					target = redactConnectionString(matches[1])
				}
				connections = append(connections, ExternalConnection{
					Type:     "object-storage",
					Service:  p.service,
					Target:   target,
					Source:   loc,
					Function: currentFunc,
				})
			}

			// gRPC
			for _, p := range pyGRPCPatterns {
				matches := p.FindStringSubmatch(stripped)
				if matches == nil {
					continue
				}
				target := ""
				if len(matches) > 1 {
					target = redactTarget(matches[1])
				}
				connections = append(connections, ExternalConnection{
					Type:     "grpc",
					Service:  "grpc",
					Target:   target,
					Source:   loc,
					Function: currentFunc,
				})
			}

			// Messaging
			for _, p := range pyMsgPatterns {
				matches := p.re.FindStringSubmatch(stripped)
				if matches == nil {
					continue
				}
				target := ""
				if len(matches) > 1 && strings.Contains(matches[1], "://") {
					target = redactConnectionString(matches[1])
				}
				connections = append(connections, ExternalConnection{
					Type:     "messaging",
					Service:  p.service,
					Target:   target,
					Source:   loc,
					Function: currentFunc,
				})
			}

			// HTTP clients
			for _, p := range pyHTTPPatterns {
				if p.re.MatchString(stripped) {
					connections = append(connections, ExternalConnection{
						Type:     "api",
						Service:  p.service,
						Target:   "",
						Source:   loc,
						Function: currentFunc,
					})
				}
			}

			// LLM/ML SDK clients
			for _, p := range pySDKPatterns {
				if p.re.MatchString(stripped) {
					connections = append(connections, ExternalConnection{
						Type:     "api",
						Service:  p.service,
						Target:   "",
						Source:   loc,
						Function: currentFunc,
					})
				}
			}
		}
	}

	return connections
}

func leadingSpaces(line string) int {
	for i, ch := range line {
		if ch != ' ' && ch != '\t' {
			return i
		}
	}
	return len(line)
}
