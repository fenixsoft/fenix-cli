package docker

import (
	"context"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"time"

	"docker.io/go-docker"
	"github.com/c-bata/go-prompt"
)

var dockerClient *docker.Client
var DockerRuntime *environments.Runtime

func RegisterEnv() (*environments.Runtime, error) {
	dockerClient, _ = docker.NewEnvClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := dockerClient.Ping(ctx); err != nil {
		return nil, err
	}

	DockerRuntime = &environments.Runtime{
		Prefix:         "docker",
		Completer:      completer,
		Commands:       ExtraCommands,
		MainSuggestion: MajorCommands,
		Executor: environments.GetDefaultExecutor("docker", func() {
			lastQueryResult = []prompt.Suggest{}
		}),
	}
	return DockerRuntime, nil
}
