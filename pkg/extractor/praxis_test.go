package extractor

import (
	"os"
	"path/filepath"
	"testing"
)

// samplePraxisConfigmap mirrors opendatahub-io/praxis-extproc
// deploy/base/config/configmap.yaml.
const samplePraxisConfigmap = `apiVersion: v1
kind: ConfigMap
metadata:
  name: payload-processing-plugins
data:
  pre-extproc.yaml: |
    server:
      grpc_address: "0.0.0.0:9004"
      health_address: "0.0.0.0:50052"
      metrics_address: "0.0.0.0:9090"
      tls:
        mode: self_signed

    filter_chains:
      - name: bbr
        filters:
          - filter: json_body
            conditions:
              - unless:
                  headers_present: ["X-Gateway-Model-Name"]
            request_extract:
              - pointer: /model
                header: X-Gateway-Model-Name
            on_invalid: continue
          - filter: llmisvc_model_provider_resolver
            header: X-Gateway-Model-Name

    insecure_options:
      allow_unbounded_body: true

  extproc.yaml: |
    server:
      grpc_address: "0.0.0.0:9004"
      health_address: "0.0.0.0:50052"
      metrics_address: "0.0.0.0:9090"
      tls:
        mode: self_signed

    filter_chains:
      - name: ipp
        filters:
          - filter: headers
            conditions:
              - unless:
                  path_prefix: /maas-api
            request_remove:
              - authorization
              - x-api-key
          - filter: request_id

    insecure_options:
      allow_unbounded_body: true
`

// writeConfigmap writes a ConfigMap fixture at the conventional deploy/base
// path so ConfigMap discovery patterns pick it up.
func writeConfigmap(t *testing.T, dir, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Join(dir, "deploy", "base", "config"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "deploy", "base", "config", "configmap.yaml"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

// TestExtractPraxisPipelines verifies extraction of both pipelines from a
// real praxis-extproc ConfigMap fixture, including ordering, server fields,
// chain and filter order, and insecure options.
func TestExtractPraxisPipelines(t *testing.T) {
	dir := t.TempDir()
	writeConfigmap(t, dir, samplePraxisConfigmap)

	pipelines := extractPraxisPipelines(dir)
	if len(pipelines) != 2 {
		t.Fatalf("expected 2 pipelines, got %d: %+v", len(pipelines), pipelines)
	}

	// Sorted by data key within the ConfigMap.
	extproc, pre := pipelines[0], pipelines[1]
	if extproc.Name != "extproc.yaml" || pre.Name != "pre-extproc.yaml" {
		t.Fatalf("unexpected pipeline order: %+v", pipelines)
	}
	for _, p := range pipelines {
		if p.ConfigMap != "payload-processing-plugins" {
			t.Errorf("%s configmap: got %q", p.Name, p.ConfigMap)
		}
		if p.Source != filepath.Join("deploy", "base", "config", "configmap.yaml") {
			t.Errorf("%s source: got %q", p.Name, p.Source)
		}
		if p.Server == nil {
			t.Fatalf("%s: missing server", p.Name)
		}
		if p.Server.GRPCAddress != "0.0.0.0:9004" || p.Server.HealthAddress != "0.0.0.0:50052" ||
			p.Server.MetricsAddress != "0.0.0.0:9090" || p.Server.TLSMode != "self_signed" {
			t.Errorf("%s server: got %+v", p.Name, p.Server)
		}
		if len(p.InsecureOptions) != 1 || p.InsecureOptions[0] != "allow_unbounded_body" {
			t.Errorf("%s insecure options: got %v", p.Name, p.InsecureOptions)
		}
	}

	// Ordered filter chains with ordered filters.
	if len(extproc.Chains) != 1 || extproc.Chains[0].Name != "ipp" {
		t.Fatalf("extproc chains: got %+v", extproc.Chains)
	}
	extFilters := extproc.Chains[0].Filters
	if len(extFilters) != 2 || extFilters[0].Filter != "headers" || extFilters[1].Filter != "request_id" {
		t.Errorf("extproc ipp filters: got %+v", extFilters)
	}
	if len(pre.Chains) != 1 || pre.Chains[0].Name != "bbr" {
		t.Fatalf("pre chains: got %+v", pre.Chains)
	}
	preFilters := pre.Chains[0].Filters
	if len(preFilters) != 2 || preFilters[0].Filter != "json_body" || preFilters[1].Filter != "llmisvc_model_provider_resolver" {
		t.Errorf("pre bbr filters: got %+v", preFilters)
	}
}

// TestExtractPraxisPipelinesNonPraxisConfigmap verifies that ConfigMap data
// values without the Praxis shape produce no pipelines.
func TestExtractPraxisPipelinesNonPraxisConfigmap(t *testing.T) {
	dir := t.TempDir()
	writeConfigmap(t, dir, `apiVersion: v1
kind: ConfigMap
metadata:
  name: plain-config
data:
  app.properties: |
    timeout=30
  not-praxis.yaml: |
    logger:
      level: info
    receivers:
      - name: otlp
`)

	if pipelines := extractPraxisPipelines(dir); len(pipelines) != 0 {
		t.Errorf("expected no pipelines from non-Praxis ConfigMap, got %+v", pipelines)
	}
}

// TestExtractPraxisPipelinesChainsOnly verifies that a pipeline with
// filter_chains but no server block is still recognized.
func TestExtractPraxisPipelinesChainsOnly(t *testing.T) {
	// A pipeline with filter_chains but no server block is still recognized.
	dir := t.TempDir()
	writeConfigmap(t, dir, `apiVersion: v1
kind: ConfigMap
metadata:
  name: chains-only
data:
  pipeline.yaml: |
    filter_chains:
      - name: solo
        filters:
          - filter: request_id
`)

	pipelines := extractPraxisPipelines(dir)
	if len(pipelines) != 1 {
		t.Fatalf("expected 1 pipeline, got %d", len(pipelines))
	}
	p := pipelines[0]
	if p.Server != nil {
		t.Errorf("expected no server, got %+v", p.Server)
	}
	if len(p.Chains) != 1 || p.Chains[0].Name != "solo" || len(p.Chains[0].Filters) != 1 {
		t.Errorf("chains: got %+v", p.Chains)
	}
}

// TestExtractPraxisPipelinesDisabledInsecureOptions verifies that only
// enabled insecure options are surfaced.
func TestExtractPraxisPipelinesDisabledInsecureOptions(t *testing.T) {
	// Only enabled insecure options are surfaced.
	dir := t.TempDir()
	writeConfigmap(t, dir, `apiVersion: v1
kind: ConfigMap
metadata:
  name: opts
data:
  pipeline.yaml: |
    server:
      grpc_address: "0.0.0.0:9004"
    insecure_options:
      allow_unbounded_body: false
`)

	pipelines := extractPraxisPipelines(dir)
	if len(pipelines) != 1 {
		t.Fatalf("expected 1 pipeline, got %d", len(pipelines))
	}
	if len(pipelines[0].InsecureOptions) != 0 {
		t.Errorf("expected no insecure options, got %v", pipelines[0].InsecureOptions)
	}
}

// TestExtractPraxisPipelinesServerWithoutGRPCAddress verifies that a
// server map without grpc_address, or a field name matched only in a
// comment, is not accepted as a pipeline.
func TestExtractPraxisPipelinesServerWithoutGRPCAddress(t *testing.T) {
	// A server map without grpc_address (and no filter_chains) must not be
	// accepted, even when the substring pre-check matches a comment or
	// another field.
	dir := t.TempDir()
	writeConfigmap(t, dir, `apiVersion: v1
kind: ConfigMap
metadata:
  name: not-praxis
data:
  mention.yaml: |
    # grpc_address appears only in this comment
    logger:
      level: info
  serverish.yaml: |
    server:
      mode: opaque
      grpc_address_note: not the real field
`)

	if pipelines := extractPraxisPipelines(dir); len(pipelines) != 0 {
		t.Errorf("expected no pipelines, got %+v", pipelines)
	}
}

// TestPraxisPipelinesDataCoverage verifies the praxis_pipelines coverage
// classification.
func TestPraxisPipelinesDataCoverage(t *testing.T) {
	dir := t.TempDir()
	writeConfigmap(t, dir, samplePraxisConfigmap)

	arch := &ComponentArchitecture{
		PraxisPipelines: extractPraxisPipelines(dir),
	}
	cov := computeDataCoverage(arch)
	if cov["praxis_pipelines"] != "sparse" {
		t.Errorf("data_coverage[praxis_pipelines]: got %q, want sparse", cov["praxis_pipelines"])
	}
}
