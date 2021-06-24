package docker

import (
	"context"
	"docker.io/go-docker/api/types"
	"fmt"
	"github.com/c-bata/go-prompt"
	"reflect"
	"strings"
)

const (
	Orchestrator = "orchestrator"
	Image        = "image"
	Container    = "container"
	Port         = "port"
	DockerType   = "docker_type"
)

func provideContainerSuggestion(args ...string) []prompt.Suggest {
	var suggestions []prompt.Suggest
	ctx := context.Background()
	cList, _ := dockerClient.ContainerList(ctx, types.ContainerListOptions{All: true})

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

func providePortSuggestion(args ...string) []prompt.Suggest {
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

func provideImagesSuggestion(arg ...string) []prompt.Suggest {
	images, _ := dockerClient.ImageList(context.Background(), types.ImageListOptions{All: true})
	var suggestions []prompt.Suggest

	for _, image := range images {
		ins, _, _ := dockerClient.ImageInspectWithRaw(context.Background(), image.ID)
		suggestions = append(suggestions, prompt.Suggest{Text: getImageID(image), Description: getImageDescription(image, ins)})
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
	if id == "" {
		id = image.ID[7:19]
	}
	return fmt.Sprintf("%-64v", id)
}

func getImageDescription(image types.ImageSummary, inspect types.ImageInspect) string {
	size := float64(inspect.Size) / (1024 * 1024)
	desc := fmt.Sprintf("%v | %8.2f MB | %v", inspect.ID[7:19], size, inspect.Created[0:10])
	return desc
}
