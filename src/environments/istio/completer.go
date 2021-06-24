package istio

import (
	"errors"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes/kube"
	"github.com/fenixsoft/fenix-cli/src/suggestions"
)

func New() (*environments.RuntimeCompleter, error) {
	// only support Istio over Kubernetes
	// so that must already installed Kubernetes Env
	if _, err := kube.NewCompleter(); err != nil {
		return nil, err
	}
	// check if istioctl is valid in PATH
	code, msg := environments.ExecuteAndGetResult("istioctl", "version")
	if code != 0 {
		return nil, errors.New(msg)
	}

	completer := &environments.RuntimeCompleter{
		GenericCompleter: suggestions.NewGenericCompleter(arguments, options, func(c *suggestions.GenericCompleter) {
			c.SuggestionProviders.Add(suggestions.Argument, suggestions.BuildStaticProvider(c.Arguments, suggestions.LengthFilter))
			c.SuggestionProviders.Add(suggestions.Option, suggestions.BuildStaticProvider(c.Options, suggestions.BestEffortFilter))
		}),
	}

	return completer, nil
}
