package environments

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	prompt "github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/internal/template"
	"os"
)

type Runtime struct {
	Prefix         string
	Completer      prompt.Completer
	Executor       prompt.Executor
	ExitChecker    prompt.ExitChecker
	LivePrefix     func() (prefix string, useLivePrefix bool)
	Commands       []Command
	MainSuggestion []prompt.Suggest
}

type Register func() (*Runtime, error)

type Environment string

const (
	Kubernetes Environment = "Kubernetes"
	Docker     Environment = "Docker"
	Rancher    Environment = "Rancher"
	Istio      Environment = "Istio"
)

var Environments = map[Environment]*Runtime{}

var Registers = map[Environment]Register{}

var ActiveEnvironment Environment

// Select one of installed environment
func SelectCurrentEnvironment() {
	if len(Environments) == 0 {
		// if there is no env, exit programme
		fmt.Printf("There is no Kubernete/Docker/Istio/Rancher envirnoment")
		os.Exit(0)
	} else if len(Environments) == 1 {
		// if there is only one env, just make it default, without select menu
		for k, _ := range Environments {
			ActiveEnvironment = k
		}
	} else {
		// let user to select one
		var result string
		var options []string
		for k, _ := range Environments {
			options = append(options, string(k))
		}
		var pt = &survey.Select{
			Message: "Multiple environments are detected. You want to operate",
			Options: options,
		}
		if err := survey.AskOne(pt, &result); err == nil {
			ActiveEnvironment = Environment(result)
		} else {
			fmt.Printf("User interruptd, bye!")
			os.Exit(0)
		}
	}
}

// Check and initialize environments
func Initialize() {
	template.SetSurveyTemplate()

	extraCmds := DefaultCommands
	for k, v := range Registers {
		if env, err := v(); err == nil {
			Environments[k] = env
			extraCmds = append(extraCmds, env.Commands...)
		}
	}

	for _, v := range extraCmds {
		var envs []Environment
		if len(v.Environments) == 0 {
			envs = []Environment{Docker, Kubernetes, Istio, Rancher}
		} else {
			envs = v.Environments
		}
		for _, e := range envs {
			if env, ok := Environments[e]; ok {
				env.MainSuggestion = append(env.MainSuggestion, prompt.Suggest{Text: v.Text, Description: v.Description})
			}
		}
	}
}

func GetActive() *Runtime {
	return Environments[ActiveEnvironment]
}

func LivePrefix() (prefix string, useLivePrefix bool) {
	act := GetActive()
	if act.LivePrefix == nil {
		return defaultLivePrefix()
	} else {
		return act.LivePrefix()
	}
}

func defaultLivePrefix() (prefix string, useLivePrefix bool) {
	return GetActive().Prefix + " ", true
}

// Return active completer
func GetActiveCompleter() prompt.Completer {
	return func(document prompt.Document) []prompt.Suggest {
		return GetActive().Completer(document)
	}
}

// Return active executor
func GetActiveExecutor() prompt.Executor {
	return func(cmd string) {
		GetActive().Executor(cmd)
	}
}

func GetKubernetes() *Runtime { return Environments[Kubernetes] }
func GetDocker() *Runtime     { return Environments[Docker] }
func GetIstio() *Runtime      { return Environments[Istio] }
func GetRancher() *Runtime    { return Environments[Rancher] }
