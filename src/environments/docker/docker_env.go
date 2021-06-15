package docker

import (
	"context"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"time"

	"docker.io/go-docker"
	"github.com/c-bata/go-prompt"
)

var dockerClient *docker.Client

func RegisterEnv() (*environments.Prompt, error) {
	dockerClient, _ = docker.NewEnvClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := dockerClient.Ping(ctx); err != nil {
		return nil, err
	}

	// register global commands
	for _, cmd := range environments.DefaultCommands {
		suggest := prompt.Suggest{
			Text:        cmd.Text,
			Description: cmd.Description,
		}
		MajorCommands = append(MajorCommands, suggest)
	}

	return &environments.Prompt{
		Prefix:    "docker",
		Completer: completer,
		Executor: environments.GetDefaultExecutor("docker", func() {
			lastQueryResult = []prompt.Suggest{}
		}),
		Commands: ExtraCommands,
	}, nil
}
