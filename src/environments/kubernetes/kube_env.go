package kubernetes

import (
	"fmt"
	"github.com/fenixsoft/fenix-cli/src/internal/krew"

	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes/kube"
)

var Completer *kube.Completer

//wrapper for kubernetes completer changed (while namespace change)
func Complete(d prompt.Document) []prompt.Suggest {
	return Completer.Complete(d)
}

func RegisterEnv() (*environments.Runtime, error) {
	if c, err := kube.NewCompleter(); err != nil {
		return nil, err
	} else {
		Completer = c

		Completer.KubernetesRuntime = &environments.Runtime{
			Prefix:         "kubectl",
			Completer:      Complete,
			Executor:       environments.GetDefaultExecutor("kubectl", nil, krew.GetBinPath()...),
			Commands:       ExtraCommands,
			MainSuggestion: kube.Commands,
			LivePrefix: func() (prefix string, useLivePrefix bool) {
				return fmt.Sprintf("%v > kubectl ", c.Namespace), true
			},
		}
		return Completer.KubernetesRuntime, nil
	}
}
