package environments

import (
	"context"
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	prompt "github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/lib/go-ansi"
	"github.com/fenixsoft/fenix-cli/src/internal/template"
	"github.com/fenixsoft/fenix-cli/src/suggestions"
	"github.com/schollz/progressbar/v3"
	"os"
	"time"
)

type Runtime struct {
	Prefix     string
	Completer  *suggestions.GenericCompleter
	Executor   prompt.Executor
	Setup      func()
	LivePrefix func() (prefix string, useLivePrefix bool)
	Commands   []Command
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

	bar := progressbar.NewOptions(len(Registers)+1,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowIts(),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionSetWidth(40),
		progressbar.OptionSetDescription("[light_blue]Scanning Environment...[reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[yellow]=[reset]",
			SaucerHead:    "[yellow]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	extraCmds := DefaultCommands
	for k, v := range Registers {
		_ = bar.Add(1)
		if env, err := TimeoutRegister(v, 5*time.Second); err == nil {
			Environments[k] = env
			extraCmds = append(extraCmds, env.Commands...)
		}
		time.Sleep(100 * time.Millisecond)
	}
	_ = bar.Add(1)
	// clear entire line and ove cursor to beginning of the line 1 lines down.
	ansi.EraseInLine(2)
	ansi.CursorNextLine(1)

	for _, v := range extraCmds {
		var envs []Environment
		if len(v.Environments) == 0 {
			envs = []Environment{Docker, Kubernetes, Istio, Rancher}
		} else {
			envs = v.Environments
		}
		for _, e := range envs {
			if env, ok := Environments[e]; ok {
				c := env.Completer
				if c != nil {
					c.Arguments = append(c.Arguments, prompt.Suggest{Text: v.Text, Alias: v.Alias, Provider: v.Provider, Description: v.Description})
					c.Arguments = append(c.Arguments, v.ExtendArguments...)
					c.Options = append(c.Options, v.ExtendOptions...)
				}
			}
		}
	}

	for k := range Environments {
		c := Environments[k].Completer
		if c != nil {
			c.Setup(c)
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
		return GetActive().Completer.Complete(document)
	}
}

// Return active executor
func GetActiveExecutor() prompt.Executor {
	return func(cmd string) {
		GetActive().Executor(cmd)
	}
}

func TimeoutRegister(fn Register, timeout time.Duration) (*Runtime, error) {
	ctx := context.Background()
	done := make(chan *Runtime, 1)
	err := make(chan error, 1)

	go func(ctx context.Context) {
		if r, e := fn(); e != nil {
			err <- e
		} else {
			done <- r
		}
	}(ctx)

	select {
	case r := <-done:
		return r, nil
	case e := <-err:
		return nil, e
	case <-time.After(timeout):
		return nil, errors.New("execution timeout")
	}
}

func GetKubernetes() *Runtime { return Environments[Kubernetes] }
func GetDocker() *Runtime     { return Environments[Docker] }
func GetIstio() *Runtime      { return Environments[Istio] }
func GetRancher() *Runtime    { return Environments[Rancher] }
