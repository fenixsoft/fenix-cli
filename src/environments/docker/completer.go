package docker

import (
	"context"
	"docker.io/go-docker/api/types"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"reflect"
	"strings"
)

func completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")
	w := d.GetWordBeforeCursor()

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(w, "-") {
		return optionCompleter(args, w)
	}

	commandArgs, optionValue := excludeOptions(args)
	if optionValue {
		// when type 'get pod -o ', we don't want to complete pods. we want to type 'json' or other.
		// So we need to skip argumentCompleter.
		return optionValueCompleter(args, w)
	}

	return argumentsCompleter(commandArgs, w)
}

func containerSuggestion() []prompt.Suggest {
	return containerListCompleter(true)
}

func containerListCompleter(all bool) []prompt.Suggest {
	var suggestions []prompt.Suggest
	ctx := context.Background()
	cList, _ := dockerClient.ContainerList(ctx, types.ContainerListOptions{All: all})

	for _, container := range cList {
		suggestions = append(suggestions, prompt.Suggest{Text: container.ID[0:12], Description: getContainerDescription(container)})
	}
	return suggestions
}

func getContainerDescription(container types.Container) string {
	id := container.Image
	if strings.HasPrefix(id, "sha256") {
		id = container.Command
	} else if strings.Contains(id, "@sha256") {
		id = id[0:strings.Index(id, "@sha256")]
	}
	return fmt.Sprintf("%7v | %v", strings.ToUpper(container.State), id)
}

func portMappingSuggestion() []prompt.Suggest {
	images, _ := dockerClient.ImageList(context.Background(), types.ImageListOptions{All: true})
	var suggestions []prompt.Suggest

	exists := map[string]bool{}
	for _, image := range images {
		inspection, _, _ := dockerClient.ImageInspectWithRaw(context.Background(), image.ID)
		exposedPortKeys := reflect.ValueOf(inspection.Config.ExposedPorts).MapKeys()
		for _, exposedPort := range exposedPortKeys {
			if _, ok := exists[exposedPort.String()]; ok {
				continue
			} else {
				exists[exposedPort.String()] = true
			}
			portAndType := strings.Split(exposedPort.String(), "/")
			port := portAndType[0]
			portType := portAndType[1]
			suggestions = append(suggestions, prompt.Suggest{Text: fmt.Sprintf("%s:%s/%s", port, port, portType), Description: getImageID(image)})
		}
	}
	return suggestions
}

func getImageID(image types.ImageSummary) string {
	id := ""
	if len(image.RepoTags) > 0 && image.RepoTags[0] != "" {
		id = image.RepoTags[0]
	} else if len(image.RepoDigests) > 0 && image.RepoDigests[0] != "" {
		id = image.RepoDigests[0]
	}
	if strings.Contains(id, "@sha256") {
		id = id[0:strings.Index(id, "@sha256")]
	}
	if id == "" || id == "<none>:<none>" {
		return image.ID[7:19]
	} else {
		return id
	}
}

func getImageDescription(image types.ImageInspect) string {
	desc := ""
	for _, v := range image.Config.Entrypoint {
		desc += v + " "
	}
	return desc
}

func imagesSuggestion() []prompt.Suggest {
	images, _ := dockerClient.ImageList(context.Background(), types.ImageListOptions{All: true})
	var suggestions []prompt.Suggest

	for _, image := range images {
		ins, _, _ := dockerClient.ImageInspectWithRaw(context.Background(), image.ID)
		suggestions = append(suggestions, prompt.Suggest{Text: getImageID(image), Description: getImageDescription(ins)})
	}
	return suggestions
}

func optionCompleter(args []string, currentArg string) []prompt.Suggest {
	l := len(args)
	if l <= 1 {
		return prompt.FilterContains(optionsDocker["global"], currentArg, true)
	} else {
		switch args[0] {
		case "app", "builder", "buildx", "checkpoint", "cluster", "compose", "config",
			"container", "context", "image", "manifest", "network", "node", "plugin",
			"secret", "service", "stack", "swarm", "system", "trust", "volume", "x-batch":
			return prompt.FilterHasPrefix(optionsDocker[args[0]+" "+args[1]], currentArg, true)
		default:
			return prompt.FilterHasPrefix(optionsDocker[args[0]], currentArg, true)
		}
	}
}

func excludeOptions(args []string) ([]string, bool) {
	l := len(args)
	if l == 0 {
		return nil, false
	}
	filtered := make([]string, 0, l)

	var optionValueFlag bool
	for i := 0; i < len(args)-1; i++ {
		if optionValueFlag {
			optionValueFlag = false
			continue
		}

		// ignore first or continuous blank spaces
		if (i == 0 && args[0] == "") || (args[i] == "" && args[i-1] == "") {
			continue
		}

		if strings.HasPrefix(args[i], "-") {
			optionValueFlag = true
			// we can specify option value like '-o=json'
			if strings.Contains(args[i], "=") {
				optionValueFlag = false
			}
			// these args are single switch arg
			for _, s := range []string{
				"--no-stdin", "--no-cache", "--compress", "--force-rm", "--rm", "--pull", "--quiet", "-q", "--squash",
				"--pause", "-p", "--archive", "--help", "--restart", "--pull", "--interactive", "-i", "--no-trunc",
				"--human", "-H", "--all", "--digests", "--password-stdin", "--details", "--latest", "-l", "--force",
				"-f", "--no-prune", "--publish-all", "-P", "-t", "-it", "--detach", "-d",
			} {
				if strings.HasPrefix(args[i], s) {
					if s == "-f" || s == "-p" {
						// check -f/-p meat --file/--publish or --force/--pause
						switch args[0] {
						case "builder", "buildx", "build", "compose", "image", "run", "create":
							// --file/--publish
							optionValueFlag = true
						default:
							// --force/--pause
							optionValueFlag = false
						}
					} else {
						optionValueFlag = false
					}
				}
			}
			continue
		}

		filtered = append(filtered, args[i])
	}
	filtered = append(filtered, args[len(args)-1])
	return filtered, optionValueFlag
}

func optionValueCompleter(args []string, currentArg string) []prompt.Suggest {
	l := len(args)
	switch args[0] {
	case "builder", "buildx", "build", "compose", "image":
		if l >= 2 && (args[l-2] == "-f" || args[l-2] == "--file") {
			return environments.GetPathSuggestion(currentArg)
		}
	case "run", "create":
		if l > 2 && (args[l-2] == "-p" || args[l-2] == "--publish") {
			return prompt.FilterFuzzy(portMappingSuggestion(), currentArg, true)
		}
	}
	return []prompt.Suggest{}
}

func argumentsCompleter(args []string, currentArg string) []prompt.Suggest {
	l := len(args)
	if l == 0 {
		return DockerRuntime.MainSuggestion
	} else if l == 1 {
		return prompt.FilterHasPrefix(DockerRuntime.MainSuggestion, args[0], true)
	}

	switch args[0] {
	case "exec", "stop", "restart", "kill", "pause", "unpause", "update",
		"wait", "logs", "attach", "top", "port", "rename", "stats":
		if l == 2 {
			return prompt.FilterDescContains(containerListCompleter(false), currentArg, true)
		} else {
			return []prompt.Suggest{}
		}
	case "start", "export", "commit", "rm", "diff", "inspect":
		if l == 2 {
			return prompt.FilterDescContains(containerListCompleter(true), currentArg, true)
		} else {
			return []prompt.Suggest{}
		}
	case "rmi", "tag", "history", "push", "save", "run", "create":
		if l == 2 {
			return prompt.FilterFuzzy(imagesSuggestion(), currentArg, true)
		} else {
			return []prompt.Suggest{}
		}
	case "pull":
		// total args > 2 or chars in current arg (should be image name) < 2, return empty
		if l > 2 || len(currentArg) <= 2 {
			return []prompt.Suggest{}
		} else {
			imageKeyword = currentArg
			imageQuery.Trigger()
			return lastQueryResult
		}
	case "app", "builder", "buildx", "checkpoint", "cluster", "compose", "config",
		"container", "context", "image", "manifest", "network", "node", "plugin",
		"secret", "service", "stack", "swarm", "system", "trust", "volume", "x-batch":
		if l == 2 {
			return prompt.FilterHasPrefix(subCommands[args[0]], currentArg, true)
		} else {
			return []prompt.Suggest{}
		}
	case "build", "import":
		if l == 2 {
			return environments.GetPathSuggestion(currentArg)
		} else {
			return []prompt.Suggest{}
		}
	default:
		return []prompt.Suggest{}
	}
}
