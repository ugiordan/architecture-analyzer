package extractor

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var dockerfilePatterns = []string{
	"Dockerfile",
	"Dockerfile.*",
	"*.Dockerfile",
	"Containerfile",
	"Containerfile.*",
	"*.Containerfile",
	"**/Dockerfile",
	"**/Dockerfile.*",
	"**/*.Dockerfile",
	"**/Containerfile",
	"**/Containerfile.*",
	"**/*.Containerfile",
}

var (
	fromRE       = regexp.MustCompile(`(?i)^FROM\s+(?:--platform=(\S+)\s+)?(\S+)(?:\s+AS\s+(\S+))?`)
	userRE       = regexp.MustCompile(`^USER\s+(\S+)`)
	exposeRE     = regexp.MustCompile(`^EXPOSE\s+(.*)`)
	argRE        = regexp.MustCompile(`^ARG\s+(\w+)(?:=(.*))?`)
	envRE        = regexp.MustCompile(`^ENV\s+(\w+)(?:=|\s+)(.*)`)
	targetArchRE = regexp.MustCompile(`TARGETARCH|TARGETOS|TARGETPLATFORM`)
	copyRE       = regexp.MustCompile(`(?i)^COPY\s+((?:--\w+(?:=\S+)?\s+)*)(.+)`)
	addRE        = regexp.MustCompile(`(?i)^ADD\s+((?:--\w+(?:=\S+)?\s+)*)(.+)`)
	copyFromRE   = regexp.MustCompile(`--from=(\S+)`)

	runRE             = regexp.MustCompile(`(?i)^RUN\s+(.+)`)
	goBuildRE         = regexp.MustCompile(`go\s+build\b`)
	goBuildOutputRE   = regexp.MustCompile(`-o\s+(\S+)`)
	goBuildEntryRE    = regexp.MustCompile(`\s((?:\./)?(?:[\w./-]+\.go|[\w./-]+/\.\.\.|\.(?:/[\w./-]+)+/?))$`)
	npmBuildRE        = regexp.MustCompile(`(?:npm\s+run|yarn\s+run|yarn|pnpm(?:\s+run)?)\s+(build\S*)`)
	pipInstallRE      = regexp.MustCompile(`pip3?\s+install\b`)
	pipRequirementsRE = regexp.MustCompile(`-r\s+(\S+)`)
	makeTargetRE      = regexp.MustCompile(`\bmake\s+([a-zA-Z0-9._-]+)`)
)

// FIPS-related environment variables and build args.
var fipsIndicators = []string{
	"GOEXPERIMENT=boringcrypto",
	"FIPS_MODE",
	"OPENSSL_FIPS",
	"FIPS_ENABLED",
}

// extractDockerfiles scans Dockerfiles/Containerfiles for metadata including
// base images, user directives, exposed ports, architectures, and FIPS flags.
func extractDockerfiles(repoPath string) []DockerfileInfo {
	files := findFiles(repoPath, dockerfilePatterns)
	var dockerfiles []DockerfileInfo

	for _, fpath := range files {
		data, err := os.ReadFile(fpath)
		if err != nil {
			log.Printf("warning: skipping %s: %v", fpath, err)
			continue
		}

		lines := joinContinuationLines(strings.Split(string(data), "\n"))
		var fromImages []string
		user := ""
		var exposedPorts []int
		var issues []string
		archSet := make(map[string]bool)
		fipsEnabled := false
		buildArgs := make(map[string]string)
		content := string(data)
		type stagedCopy struct {
			CopyInstruction
			stage int
		}
		stageAlias := make(map[string]int)
		stageParent := make(map[int]int)
		stageIndex := -1
		var stagedCopies []stagedCopy
		stageBuildCmds := make(map[int][]BuildCommand)

		for _, line := range lines {
			stripped := strings.TrimSpace(line)
			if stripped == "" || strings.HasPrefix(stripped, "#") {
				continue
			}

			// FROM instruction with optional --platform
			if match := fromRE.FindStringSubmatch(stripped); match != nil {
				platform := match[1]
				image := match[2]
				alias := match[3]
				fromImages = append(fromImages, image)

				recordStage(&stageIndex, stageAlias, stageParent, image, alias)

				if platform != "" {
					// --platform=linux/amd64 or $TARGETPLATFORM
					if strings.Contains(platform, "$") {
						// Variable reference, will check for TARGETARCH usage
					} else {
						parts := strings.Split(platform, "/")
						if len(parts) >= 2 {
							archSet[parts[1]] = true
						}
					}
				}

				if strings.HasSuffix(image, ":latest") ||
					(!strings.Contains(image, ":") && !strings.Contains(image, "@")) {
					issues = append(issues, fmt.Sprintf("Unpinned base image: %s", image))
				}
			}

			// USER instruction
			if match := userRE.FindStringSubmatch(stripped); match != nil {
				user = match[1]
			}

			// EXPOSE instruction
			if match := exposeRE.FindStringSubmatch(stripped); match != nil {
				for _, part := range strings.Fields(match[1]) {
					portStr := strings.SplitN(part, "/", 2)[0]
					if port, err := strconv.Atoi(portStr); err == nil {
						exposedPorts = append(exposedPorts, port)
					}
				}
			}

			// ARG instruction
			if match := argRE.FindStringSubmatch(stripped); match != nil {
				argName := match[1]
				argVal := ""
				if len(match) > 2 {
					argVal = strings.TrimSpace(match[2])
				}
				// Track security-relevant build args
				if isSecurityRelevantArg(argName) {
					buildArgs[argName] = argVal
				}
			}

			// ENV instruction
			if match := envRE.FindStringSubmatch(stripped); match != nil {
				envName := match[1]
				envVal := strings.TrimSpace(match[2])
				for _, indicator := range fipsIndicators {
					if strings.Contains(envName+"="+envVal, indicator) {
						fipsEnabled = true
					}
				}
			}

			// COPY instruction
			if match := copyRE.FindStringSubmatch(stripped); match != nil {
				ci := parseCopyAdd(match[1], match[2])
				stagedCopies = append(stagedCopies, stagedCopy{CopyInstruction: ci, stage: stageIndex})
			}

			// ADD instruction
			if match := addRE.FindStringSubmatch(stripped); match != nil {
				ci := parseCopyAdd(match[1], match[2])
				stagedCopies = append(stagedCopies, stagedCopy{CopyInstruction: ci, stage: stageIndex})
			}

			// RUN instruction — extract build tool invocations
			if match := runRE.FindStringSubmatch(stripped); match != nil {
				stageBuildCmds[stageIndex] = append(stageBuildCmds[stageIndex], parseRunCommand(match[1])...)
			}
		}

		// Check for TARGETARCH usage (multi-arch build)
		if targetArchRE.MatchString(content) {
			archSet["multi-arch"] = true
		}

		// Check for root user
		if user == "root" || user == "0" {
			issues = append(issues, "Container runs as root user")
		} else if user == "" {
			issues = append(issues, "No USER directive found (defaults to root)")
		}

		// Check for FIPS in build args
		for argName, argVal := range buildArgs {
			for _, indicator := range fipsIndicators {
				if strings.Contains(argName+"="+argVal, indicator) {
					fipsEnabled = true
				}
			}
		}

		// Build per-stage host sources map
		stageHostSources := make(map[int][]string)
		for _, sc := range stagedCopies {
			if sc.FromStage == "" && !sc.IsURL {
				stageHostSources[sc.stage] = append(stageHostSources[sc.stage], sc.Sources...)
			}
		}

		// Keep only final-stage copies, trace original sources for --from references
		var finalCopies []CopyInstruction
		for _, sc := range stagedCopies {
			if sc.stage != stageIndex {
				continue
			}
			ci := sc.CopyInstruction
			if ci.FromStage != "" {
				// External stage refs (e.g., --from=external-image:tag) are not resolved
				// to host sources; OriginalSources remains nil for those cases.
				if targetIdx, ok := stageAlias[ci.FromStage]; ok {
					ci.OriginalSources = collectStageHostSources(targetIdx, stageHostSources, stageParent)
				}
			}
			finalCopies = append(finalCopies, ci)
		}

		stages := len(fromImages)
		baseImage := ""
		var buildStageImages []string
		if len(fromImages) > 0 {
			baseImage = fromImages[len(fromImages)-1] // runtime image is the last FROM stage
			// Capture all non-final stage images (build stages)
			if len(fromImages) > 1 {
				buildStageImages = fromImages[:len(fromImages)-1]
			}
		}

		if exposedPorts == nil {
			exposedPorts = []int{}
		}
		if issues == nil {
			issues = []string{}
		}

		var architectures []string
		for arch := range archSet {
			architectures = append(architectures, arch)
		}
		sort.Strings(architectures)
		if architectures == nil {
			architectures = []string{}
		}

		if len(buildArgs) == 0 {
			buildArgs = nil
		}

		// Collect build commands from all stages that contribute to the final image
		var buildCmds []BuildCommand
		referencedStages := make(map[int]bool)
		referencedStages[stageIndex] = true
		worklist := []int{stageIndex}
		for len(worklist) > 0 {
			cur := worklist[len(worklist)-1]
			worklist = worklist[:len(worklist)-1]
			// Follow parent lineage
			if parent, ok := stageParent[cur]; ok && !referencedStages[parent] {
				referencedStages[parent] = true
				worklist = append(worklist, parent)
			}
			// Follow transitive COPY --from references from this stage
			for _, sc := range stagedCopies {
				if sc.stage == cur && sc.FromStage != "" {
					if idx, ok := stageAlias[sc.FromStage]; ok && !referencedStages[idx] {
						referencedStages[idx] = true
						worklist = append(worklist, idx)
					}
				}
			}
		}
		for si := 0; si <= stageIndex; si++ {
			if referencedStages[si] {
				buildCmds = append(buildCmds, stageBuildCmds[si]...)
			}
		}

		dockerfiles = append(dockerfiles, DockerfileInfo{
			Path:             relativePath(repoPath, fpath),
			BaseImage:        baseImage,
			BuildStageImages: buildStageImages,
			Stages:           stages,
			User:             user,
			ExposedPorts:     exposedPorts,
			Issues:           issues,
			Architectures:    architectures,
			FIPSEnabled:      fipsEnabled,
			BuildArgs:        buildArgs,
			CopyInstructions: finalCopies,
			BuildCommands:    buildCmds,
		})
	}

	if dockerfiles == nil {
		dockerfiles = []DockerfileInfo{}
	}
	return dockerfiles
}

// joinContinuationLines merges Dockerfile lines joined with trailing backslashes
// into single logical lines.
func joinContinuationLines(lines []string) []string {
	var result []string
	var current strings.Builder
	for _, line := range lines {
		trimmed := strings.TrimRight(line, " \t")
		if strings.HasSuffix(trimmed, "\\") {
			current.WriteString(strings.TrimSuffix(trimmed, "\\"))
			current.WriteByte(' ')
		} else {
			current.WriteString(line)
			result = append(result, current.String())
			current.Reset()
		}
	}
	if current.Len() > 0 {
		result = append(result, current.String())
	}
	return result
}

func isSecurityRelevantArg(name string) bool {
	upper := strings.ToUpper(name)
	relevant := []string{
		"GOEXPERIMENT", "CGO_ENABLED", "FIPS", "TARGETARCH", "TARGETOS",
		"TARGETPLATFORM", "GO_VERSION", "GOLANG_VERSION", "USER_ID",
		"GROUP_ID", "OPENSSL",
	}
	for _, r := range relevant {
		if strings.Contains(upper, r) {
			return true
		}
	}
	return false
}

func parseCopyAdd(flags, args string) CopyInstruction {
	var ci CopyInstruction

	if m := copyFromRE.FindStringSubmatch(flags); m != nil {
		ci.FromStage = m[1]
	}

	parts := splitCopyArgs(args)
	if len(parts) >= 2 {
		ci.Sources = parts[:len(parts)-1]
		ci.Destination = parts[len(parts)-1]
	} else if len(parts) == 1 {
		ci.Sources = parts
	}

	for _, s := range ci.Sources {
		if strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") {
			ci.IsURL = true
			break
		}
	}

	return ci
}

func splitCopyArgs(args string) []string {
	args = strings.TrimSpace(args)
	if strings.HasPrefix(args, "[") {
		var parts []string
		if err := json.Unmarshal([]byte(args), &parts); err != nil {
			return nil
		}
		return parts
	}
	return strings.Fields(args)
}

// recordStage updates stage tracking maps when a FROM instruction is encountered.
// It increments the stage index, registers numeric and named aliases, and records
// parent-child relationships for multi-stage builds.
func recordStage(stageIndex *int, stageAlias map[string]int, stageParent map[int]int, image, alias string) {
	*stageIndex++
	idx := *stageIndex
	stageAlias[strconv.Itoa(idx)] = idx
	if alias != "" {
		stageAlias[alias] = idx
	}
	if parentIdx, ok := stageAlias[image]; ok {
		stageParent[idx] = parentIdx
	}
}

// parseRunCommand extracts build tool invocations from a RUN instruction body.
// A single RUN line may contain chained commands (&&); each segment is checked
// independently so multiple tools in one RUN are all captured.
func parseRunCommand(body string) []BuildCommand {
	var cmds []BuildCommand
	for _, segment := range strings.Split(body, "&&") {
		cmd := strings.TrimSpace(segment)
		if bc := parseGoBuild(cmd); bc != nil {
			cmds = append(cmds, *bc)
		} else if bc := parseNpmBuild(cmd); bc != nil {
			cmds = append(cmds, *bc)
		} else if bc := parsePipInstall(cmd); bc != nil {
			cmds = append(cmds, *bc)
		} else if bc := parseMakeBuild(cmd); bc != nil {
			cmds = append(cmds, *bc)
		}
	}
	return cmds
}

func parseGoBuild(cmd string) *BuildCommand {
	if !goBuildRE.MatchString(cmd) {
		return nil
	}
	bc := &BuildCommand{Tool: "go", Command: "build"}
	if m := goBuildOutputRE.FindStringSubmatch(cmd); m != nil {
		bc.Output = m[1]
	}
	if m := goBuildEntryRE.FindStringSubmatch(cmd); m != nil {
		bc.EntryPoint = m[1]
	}
	return bc
}

func parseNpmBuild(cmd string) *BuildCommand {
	m := npmBuildRE.FindStringSubmatch(cmd)
	if m == nil {
		return nil
	}
	tool := "npm"
	if strings.Contains(cmd, "yarn") {
		tool = "yarn"
	} else if strings.Contains(cmd, "pnpm") {
		tool = "pnpm"
	}
	return &BuildCommand{Tool: tool, Command: m[1]}
}

func parsePipInstall(cmd string) *BuildCommand {
	if !pipInstallRE.MatchString(cmd) {
		return nil
	}
	bc := &BuildCommand{Tool: "pip", Command: "install"}
	if m := pipRequirementsRE.FindStringSubmatch(cmd); m != nil {
		bc.EntryPoint = m[1]
	}
	return bc
}

func parseMakeBuild(cmd string) *BuildCommand {
	m := makeTargetRE.FindStringSubmatch(cmd)
	if m == nil {
		return nil
	}
	return &BuildCommand{Tool: "make", Command: m[1]}
}

func collectStageHostSources(idx int, hostSources map[int][]string, parents map[int]int) []string {
	seen := make(map[string]bool)
	visited := make(map[int]bool)
	var result []string
	for cur := idx; !visited[cur]; {
		visited[cur] = true
		for _, s := range hostSources[cur] {
			if !seen[s] {
				seen[s] = true
				result = append(result, s)
			}
		}
		parent, ok := parents[cur]
		if !ok {
			break
		}
		cur = parent
	}
	return result
}
