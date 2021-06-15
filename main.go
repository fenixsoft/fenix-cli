package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/docker"
	"github.com/fenixsoft/fenix-cli/src/environments/istio"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
	"os"
)

func main() {
	// init environment providers and select the active one
	environments.Registers = map[environments.Environment]environments.Register{
		environments.Docker:     docker.RegisterEnv,
		environments.Kubernetes: kubernetes.RegisterEnv,
		environments.Istio:      istio.RegisterEnv,
	}

	environments.Logo(nil, os.Stdout)

	environments.Initialize()
	environments.SelectCurrentEnvironment()

	// start go-prompt as console
	p := prompt.New(
		environments.GetActiveExecutor(),
		environments.GetActiveCompleter(),
		// register hot key for select active env
		prompt.OptionAddKeyBind(environments.BuildPromptKeyBinds()...),
		// register live prefix that will be change automatically when env changed
		prompt.OptionLivePrefix(environments.LivePrefix),
		prompt.OptionTitle("Fenix-CLI: Interactive Cloud-Native Environments Client"),
		prompt.OptionInputTextColor(prompt.Yellow),

		prompt.OptionCompletionOnDown(),
		prompt.OptionMaxSuggestion(8),
		prompt.OptionSuggestionTextColor(prompt.Black),
		prompt.OptionDescriptionTextColor(prompt.Black),
		prompt.OptionSuggestionBGColor(prompt.LightGray),
		prompt.OptionDescriptionBGColor(prompt.DarkGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.White),
		prompt.OptionSelectedDescriptionTextColor(prompt.White),
		prompt.OptionSelectedSuggestionBGColor(prompt.Blue),
		prompt.OptionSelectedDescriptionBGColor(prompt.DarkBlue),
		prompt.OptionScrollbarBGColor(prompt.LightGray),
		prompt.OptionScrollbarThumbColor(prompt.Blue),
	)

	p.Run()
}
