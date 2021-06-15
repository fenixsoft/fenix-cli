package kubernetes

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes/kube"
)

var Completer *kube.Completer

//wrapper for kubernetes completer changed (while namespace change)
func Complete(d prompt.Document) []prompt.Suggest {
	return Completer.Complete(d)
}

func RegisterEnv() (*environments.Prompt, error) {
	DisableKlog()

	if c, err := kube.NewCompleter(); err != nil {
		return nil, err
	} else {
		Completer = c

		for _, cmd := range environments.DefaultCommands {
			suggest := prompt.Suggest{
				Text:        cmd.Text,
				Description: cmd.Description,
			}
			kube.Commands = append(kube.Commands, suggest)
		}

		return &environments.Prompt{
			Prefix:    "kubectl",
			Completer: Complete,
			Executor:  environments.GetDefaultExecutor("kubectl", nil),
			LivePrefix: func() (prefix string, useLivePrefix bool) {
				return fmt.Sprintf("%v > kubectl ", c.Namespace), true
			},
			Commands: ExtraCommands,
		}, nil
	}
}
