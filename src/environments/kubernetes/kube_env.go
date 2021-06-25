package kubernetes

import (
	"fmt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/internal/krew"
)

func RegisterEnv() (*environments.Runtime, error) {
	if c, err := New(); err != nil {
		return nil, err
	} else {
		return &environments.Runtime{
			Prefix:    "kubectl",
			Completer: c,
			Executor:  environments.GetDefaultExecutor("kubectl", nil, krew.GetBinPath()...),
			Commands:  ExtraCommands,
			LivePrefix: func() (prefix string, useLivePrefix bool) {
				return fmt.Sprintf("%v > kubectl ", Client.Namespace), true
			},
		}, nil
	}
}
