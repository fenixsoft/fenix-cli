package docker

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/internal/util"
	"io"
	"strings"
)

var ExtraCommands = []environments.Command{
	{
		Text:         "x-batch",
		Description:  "Batch management of containers and images",
		Environments: []environments.Environment{environments.Docker},
		MatchFn:      environments.StartWithMatch,
		Fn: func(args []string, writer io.Writer) {
			// get all container
			if len(args) < 2 || (args[1] != "container" && args[1] != "image") {
				_, _ = writer.Write([]byte("The container or image type should be provided in the parameter\n" +
					"Try \"x-batch container\" or \"x-batch image\" again\n"))
				return
			}

			resources := map[string]struct {
				fn  func() []prompt.Suggest
				ops []string
			}{
				"container": {
					fn:  containerSuggestion,
					ops: []string{"rm", "rm -f", "start", "stop", "restart", "pause", "unpause", "kill"},
				},
				"image": {
					fn:  imagesSuggestion,
					ops: []string{"rmi", "rmi -f", "create", "run"},
				},
			}

			resource := resources[args[1]]
			var opts []string
			col := int(util.GetWindowWidth() - 15)
			for _, v := range resource.fn() {
				opts = append(opts, v.Text+" | "+util.SubString(v.Description, 0, col-len(v.Text)))
			}
			var qs = []*survey.Question{
				{
					Name: "objective",
					Prompt: &survey.MultiSelect{
						Message:  "Which resource you want to operate ?",
						Options:  opts,
						PageSize: 25,
					},
					Validate: func(val interface{}) error {
						if ans, ok := val.([]survey.OptionAnswer); !ok || len(ans) == 0 {
							return errors.New("please select a least one resource")
						}
						return nil
					},
				},
				{
					Name: "operation",
					Prompt: &survey.Select{
						Message: "What you want to do with the these resources ?",
						Options: resource.ops,
					},
				},
				{
					Name: "confirm",
					Prompt: &survey.Confirm{
						Message: "Are you sure you want to continue ?",
					},
				},
			}
			answers := struct {
				Objective []string
				Operation string
				Confirm   bool
			}{}

			err := survey.Ask(qs, &answers, survey.WithKeepFilter(true))
			if err == terminal.InterruptErr {
				_, _ = writer.Write([]byte("operation interrupted\n"))
			} else if err != nil {
				panic(err)
			}

			if answers.Confirm {
				for _, o := range answers.Objective {
					environments.Executor("docker", answers.Operation+" "+strings.TrimSpace(strings.Split(o, "|")[0]))
				}
			}
		},
	},
}
