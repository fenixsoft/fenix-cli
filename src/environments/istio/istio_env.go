package istio

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes"
)

var Completer *IstioCompleter

// MUST setup AFTER Kubernetes Env
func RegisterEnv() (*environments.Prompt, error) {
	if c, err := NewCompleter(); err != nil {
		return nil, err
	} else {
		Completer = c

		for _, cmd := range environments.DefaultCommands {
			suggest := prompt.Suggest{
				Text:        cmd.Text,
				Description: cmd.Description,
			}
			MajorCommands = append(MajorCommands, suggest)
		}

		return &environments.Prompt{
			Prefix:    "istioctl",
			Completer: c.Complete,
			Executor:  environments.GetDefaultExecutor("istioctl", nil),
			LivePrefix: func() (prefix string, useLivePrefix bool) {
				return fmt.Sprintf("%v > istioctl ", kubernetes.Completer.Namespace), true
			},
		}, nil
	}
}
