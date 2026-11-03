package extractor

import (
	"sort"
	"strings"
)

// extractPraxisPipelines scans ConfigMap definitions for embedded Praxis
// ExtProc pipeline configurations (YAML data values with a server and/or
// filter_chains structure) and captures their known structural fields.
func extractPraxisPipelines(repoPath string) []PraxisPipeline {
	var pipelines []PraxisPipeline

	for _, fpath := range findYAMLFiles(repoPath, configmapYAMLPatterns) {
		for _, doc := range parseYAMLSafe(fpath) {
			if kind, _ := doc["kind"].(string); kind != "ConfigMap" {
				continue
			}
			metadata, _ := doc["metadata"].(map[string]interface{})
			cmName := ""
			if metadata != nil {
				cmName, _ = metadata["name"].(string)
			}
			data, _ := doc["data"].(map[string]interface{})
			if len(data) == 0 {
				continue
			}

			// Deterministic iteration over data keys.
			keys := make([]string, 0, len(data))
			for k := range data {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			for _, key := range keys {
				value, _ := data[key].(string)
				pipeline, ok := parsePraxisPipeline(value)
				if !ok {
					continue
				}
				pipeline.Name = key
				pipeline.ConfigMap = cmName
				pipeline.Source = relativePath(repoPath, fpath)
				pipelines = append(pipelines, *pipeline)
			}
		}
	}

	if pipelines == nil {
		pipelines = []PraxisPipeline{}
	}
	return pipelines
}

// parsePraxisPipeline parses one ConfigMap data value as a Praxis pipeline
// configuration. The bool reports whether the value has the Praxis shape:
// a filter_chains list, or a server with a grpc_address.
func parsePraxisPipeline(value string) (*PraxisPipeline, bool) {
	if !strings.Contains(value, "filter_chains") && !strings.Contains(value, "grpc_address") {
		return nil, false
	}

	for _, m := range parseYAMLFromBytes([]byte(value)) {
		chains, hasChains := m["filter_chains"].([]interface{})
		server, hasServer := m["server"].(map[string]interface{})
		// Server-only values are accepted only with a real grpc_address; the
		// earlier substring check may have matched the field name in a
		// comment or another field.
		if !hasChains {
			if !hasServer || mapStr(server, "grpc_address") == "" {
				continue
			}
		}

		pipeline := &PraxisPipeline{}
		if hasServer {
			pipeline.Server = &PraxisServer{
				GRPCAddress:    mapStr(server, "grpc_address"),
				HealthAddress:  mapStr(server, "health_address"),
				MetricsAddress: mapStr(server, "metrics_address"),
			}
			if tls, ok := server["tls"].(map[string]interface{}); ok {
				pipeline.Server.TLSMode = mapStr(tls, "mode")
			}
		}
		for _, c := range chains {
			chain, ok := c.(map[string]interface{})
			if !ok {
				continue
			}
			fc := PraxisFilterChain{Name: mapStr(chain, "name")}
			filters, _ := chain["filters"].([]interface{})
			for _, f := range filters {
				filter, ok := f.(map[string]interface{})
				if !ok {
					continue
				}
				if name := mapStr(filter, "filter"); name != "" {
					fc.Filters = append(fc.Filters, PraxisFilter{Filter: name})
				}
			}
			pipeline.Chains = append(pipeline.Chains, fc)
		}
		if opts, ok := m["insecure_options"].(map[string]interface{}); ok {
			// Sorted flag names for deterministic output.
			names := make([]string, 0, len(opts))
			for k, v := range opts {
				if enabled, _ := v.(bool); enabled {
					names = append(names, k)
				}
			}
			sort.Strings(names)
			pipeline.InsecureOptions = names
		}
		return pipeline, true
	}
	return nil, false
}

// mapStr reads a string value from a YAML-unmarshaled map.
func mapStr(m map[string]interface{}, key string) string {
	s, _ := m[key].(string)
	return s
}
