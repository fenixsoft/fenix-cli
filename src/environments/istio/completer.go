package istio

import (
	"errors"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes"
	"github.com/fenixsoft/fenix-cli/src/suggestions"
)

func New() (*suggestions.GenericCompleter, error) {
	// only support Istio over Kubernetes
	// so that must already installed Kubernetes Env
	if _, err := kubernetes.NewClient(); err != nil {
		return nil, err
	}
	// check if istioctl is valid in PATH
	code, msg := environments.ExecuteAndGetResult("istioctl", "version")
	if code != 0 {
		return nil, errors.New(msg)
	}

	completer := suggestions.NewGenericCompleter(arguments, options, func(c *suggestions.GenericCompleter) {
		c.SuggestionProviders.Add(suggestions.Argument, suggestions.BuildStaticCompletionProvider(c.Arguments, suggestions.LengthFilter))
		c.SuggestionProviders.Add(suggestions.Option, suggestions.BuildStaticCompletionProvider(c.Options, suggestions.BestEffortFilter))
	})

	return completer, nil
}
