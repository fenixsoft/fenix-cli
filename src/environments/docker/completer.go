package docker

import (
	"context"
	"docker.io/go-docker"
	"github.com/fenixsoft/fenix-cli/src/suggestions"
	"time"
)

func New() (*suggestions.GenericCompleter, error) {
	// Determine if docker is running
	dockerClient, _ = docker.NewEnvClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := dockerClient.Ping(ctx); err != nil {
		return nil, err
	}

	return suggestions.NewGenericCompleter(arguments, options, func(c *suggestions.GenericCompleter) {
		c.SuggestionProviders.Add(suggestions.Argument, suggestions.BuildStaticCompletionProvider(c.Arguments, suggestions.LengthFilter))
		c.SuggestionProviders.Add(suggestions.Option, suggestions.BuildStaticCompletionProvider(c.Options, suggestions.BestEffortFilter))

		suggestions.RegisterSharedProvider(Image, provideImagesSuggestion)
		suggestions.RegisterSharedProvider(RemoteImage, provideRemoteImageSuggestion)
		suggestions.RegisterSharedProvider(Container, provideContainerSuggestion)
		suggestions.RegisterSharedProvider(Port, providePortSuggestion)
		suggestions.RegisterSharedProvider(Orchestrator, suggestions.BuildFixedSelectionProvider("swarm", "kubernetes", "all"))
		suggestions.RegisterSharedProvider(DockerType, suggestions.BuildFixedSelectionProvider(Image, Container))
	}), nil
}
