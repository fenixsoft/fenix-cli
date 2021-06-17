package istio

import (
	"fmt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes"
	"github.com/fenixsoft/fenix-cli/src/internal/krew"
	"github.com/pkg/errors"
)

var Completer *IstioCompleter

// MUST register AFTER Kubernetes environment
func RegisterEnv() (*environments.Runtime, error) {
	if c, err := NewCompleter(); err != nil {
		return nil, err
	} else if !krew.IsIstiocltAvailable() {
		return nil, errors.New("istio is not available")
	} else {
		Completer = c

		c.IstioRuntime = &environments.Runtime{
			Prefix:         "istioctl",
			Completer:      c.Complete,
			Executor:       environments.GetDefaultExecutor("istioctl", nil),
			MainSuggestion: MajorCommands,
			LivePrefix: func() (prefix string, useLivePrefix bool) {
				return fmt.Sprintf("%v > istioctl ", kubernetes.Completer.Namespace), true
			},
		}
		return c.IstioRuntime, nil
	}
}
