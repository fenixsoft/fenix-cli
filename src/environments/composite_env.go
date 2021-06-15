package environments

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	prompt "github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/internal/util"
	"os"
)

type Prompt struct {
	Prefix      string
	Completer   prompt.Completer
	Executor    prompt.Executor
	ExitChecker prompt.ExitChecker
	LivePrefix  func() (prefix string, useLivePrefix bool)
	Commands    []Command
}

type Register func() (*Prompt, error)

type Environment string

const (
	Kubernetes Environment = "Kubernetes"
	Docker     Environment = "Docker"
	Rancher    Environment = "Rancher"
	Istio      Environment = "Istio"
)

var Environments = map[Environment]*Prompt{}

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
		util.AssertNoError(survey.AskOne(pt, &result))
		ActiveEnvironment = Environment(result)
	}
}

// Check and initialize environments
func Initialize() {
	setSurveyTemplate()

	for k, v := range Registers {
		if env, err := v(); err == nil {
			Environments[k] = env
		}
	}
}

func GetActive() *Prompt {
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
