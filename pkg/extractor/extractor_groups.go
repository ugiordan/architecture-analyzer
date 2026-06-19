package extractor

import (
	"fmt"
	"sort"
	"strings"
)

const (
	GroupK8sCore       = "k8s_core"
	GroupCRDs          = "crds"
	GroupDocker        = "docker"
	GroupHelm          = "helm"
	GroupWebhooks      = "webhooks"
	GroupControllers   = "controllers"
	GroupKustomize     = "kustomize"
	GroupNetworking    = "networking"
	GroupObservability = "observability"
	GroupOperator      = "operator"
	GroupServing       = "serving"
	GroupAvailability  = "availability"
	GroupDependencies  = "dependencies"
	GroupPython        = "python"
	GroupComponentRefs = "component_refs"
)

type extractContext struct {
	absPath        string
	componentName  string
	modulePrefixes []string
	opts           *ExtractOptions

	// Written by observability group, read by operator enrichment in phase 2.
	statusConditionConstNames map[string]bool
	// Written by kustomize group, read by mergeKustomizeResources in phase 2.
	kustomizeResults []KustomizeBuildResult
}

type extractorGroup struct {
	name        string
	description string
	deps        []string
	needsGo     bool
	extract     func(ctx *extractContext, arch *ComponentArchitecture)
}

var allGroups = []extractorGroup{
	{
		name:        GroupK8sCore,
		description: "Core K8s resources: RBAC, Services, Deployments, NetworkPolicies, Secrets, ConfigMaps",
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.RBAC = extractRBAC(ctx.absPath)
			arch.Services = extractServices(ctx.absPath)
			arch.Deployments = extractDeployments(ctx.absPath)
			arch.NetworkPolicies = extractAllNetworkPolicies(ctx.absPath)
			arch.Secrets = extractSecrets(ctx.absPath)
			arch.ConfigMaps = extractConfigMaps(ctx.absPath)
		},
	},
	{
		name:        GroupCRDs,
		description: "CustomResourceDefinitions and API type structs",
		needsGo:     true,
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.CRDs = extractCRDs(ctx.absPath)
			arch.APITypes = extractAPITypes(ctx.absPath)
		},
	},
	{
		name:        GroupDocker,
		description: "Dockerfile analysis: base images, stages, ports, FIPS, COPY instructions",
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.Dockerfiles = extractDockerfiles(ctx.absPath)
		},
	},
	{
		name:        GroupHelm,
		description: "Helm chart metadata and values defaults",
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.Helm = extractHelm(ctx.absPath)
		},
	},
	{
		name:        GroupWebhooks,
		description: "Webhook configurations, port resolution, Go AST behavior extraction",
		needsGo:     true,
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.Webhooks = extractWebhooks(ctx.absPath)
		},
	},
	{
		name:        GroupControllers,
		description: "Controller watches, reconcile sequences, cache config, resource ops",
		deps:        []string{GroupK8sCore},
		needsGo:     true,
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.ControllerWatch = extractControllerWatches(ctx.absPath)
			arch.ReconcileSequences = extractReconcileSequences(ctx.absPath)
		},
	},
	{
		name:        GroupKustomize,
		description: "Kustomize components, overlay refs, build+merge",
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.KustomizeComponents = extractKustomizeComponents(ctx.absPath)
			arch.KustomizeOverlayRefs = extractKustomizeOverlayRefs(ctx.absPath)
			ctx.kustomizeResults = kustomizeBuildOverlays(ctx.absPath, ctx.opts.OverlayPreference)
		},
	},
	{
		name:        GroupNetworking,
		description: "Ingress routing, external connections, HTTP endpoints, runtime dependencies",
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.IngressRouting = extractIngress(ctx.absPath)
			arch.ExternalConnections = extractAllExternalConnections(ctx.absPath)
			arch.HTTPEndpoints = extractHTTPEndpoints(ctx.absPath)
			arch.RuntimeDependencies = extractRuntimeDependencies(ctx.absPath)
		},
	},
	{
		name:        GroupObservability,
		description: "Prometheus metrics, status conditions, feature gates",
		needsGo:     true,
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.PrometheusMetrics = extractPrometheusMetrics(ctx.absPath)
			arch.StatusConditions, ctx.statusConditionConstNames = extractStatusConditions(ctx.absPath)
			arch.FeatureGates = extractFeatureGates(ctx.absPath)
		},
	},
	{
		name:        GroupOperator,
		description: "Operator config, platform detection, label contracts, template files",
		deps:        []string{GroupObservability},
		needsGo:     true,
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.PlatformDetection = extractPlatformDetection(ctx.absPath)
			arch.TemplateFiles = findTemplateFiles(ctx.absPath)
			arch.LabelContracts = extractLabelContracts(ctx.absPath)
		},
	},
	{
		name:        GroupServing,
		description: "KServe/ModelMesh serving runtimes, runtime refs, resource defaults",
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.ServingRuntimes = extractServingRuntimes(ctx.absPath)
			arch.ServingRuntimeRefs = extractServingRuntimeRefs(ctx.absPath)
			arch.ResourceDefaults = extractResourceDefaults(ctx.absPath)
		},
	},
	{
		name:        GroupAvailability,
		description: "PodDisruptionBudgets, HorizontalPodAutoscalers",
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.PodDisruptionBudgets = extractPDBs(ctx.absPath)
			arch.HorizontalPodAutoscalers = extractHPAs(ctx.absPath)
		},
	},
	{
		name:        GroupDependencies,
		description: "Go modules and Python package dependencies",
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.Dependencies = extractDependencies(ctx.absPath, ctx.modulePrefixes)
		},
	},
	{
		name:        GroupPython,
		description: "Python K8s API calls and port detection",
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.PythonK8sCalls = extractPythonK8sCalls(ctx.absPath)
		},
	},
	{
		name:        GroupComponentRefs,
		description: "Cross-component references",
		extract: func(ctx *extractContext, arch *ComponentArchitecture) {
			arch.ComponentRefs = extractComponentRefs(ctx.absPath, ctx.componentName, ctx.opts.KnownComponents)
		},
	},
}

func resolveExtractorGroups(requested []string) (map[string]bool, error) {
	index := groupIndex()
	if len(requested) == 0 {
		active := make(map[string]bool, len(allGroups))
		for _, g := range allGroups {
			active[g.name] = true
		}
		return active, nil
	}

	active := make(map[string]bool)
	for _, name := range requested {
		name = strings.TrimSpace(name)
		if _, ok := index[name]; !ok {
			return nil, fmt.Errorf("unknown extractor group %q (available: %s)", name, strings.Join(ExtractorGroupNames(), ", "))
		}
		active[name] = true
	}

	for name := range active {
		if err := resolveDeps(name, active, index); err != nil {
			return nil, err
		}
	}

	return active, nil
}

func resolveDeps(name string, active map[string]bool, index map[string]*extractorGroup) error {
	grp := index[name]
	if grp == nil {
		return nil
	}
	for _, dep := range grp.deps {
		if active[dep] {
			continue
		}
		if _, ok := index[dep]; !ok {
			return fmt.Errorf("extractor group %q depends on unknown group %q", name, dep)
		}
		active[dep] = true
		if err := resolveDeps(dep, active, index); err != nil {
			return err
		}
	}
	return nil
}

func groupIndex() map[string]*extractorGroup {
	idx := make(map[string]*extractorGroup, len(allGroups))
	for i := range allGroups {
		idx[allGroups[i].name] = &allGroups[i]
	}
	return idx
}

func shouldRun(active map[string]bool, group string) bool {
	return active[group]
}

// ExtractorGroupNames returns sorted group names for CLI help.
func ExtractorGroupNames() []string {
	names := make([]string, 0, len(allGroups))
	for _, g := range allGroups {
		names = append(names, g.name)
	}
	sort.Strings(names)
	return names
}

func needsGoPackages(active map[string]bool) bool {
	for _, g := range allGroups {
		if g.needsGo && active[g.name] {
			return true
		}
	}
	return false
}
