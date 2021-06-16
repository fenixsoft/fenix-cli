package environments

import (
	"bytes"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/src/internal/util"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Command struct {
	Text         string
	HokKey       prompt.Key
	Environments []Environment
	Description  string
	MatchFn      func(string, string) bool
	Fn           func([]string, io.Writer)
}

func checkEnvironments(envs []Environment) bool {
	if len(envs) == 0 {
		return true
	} else {
		for _, env := range envs {
			if env == ActiveEnvironment {
				return true
			}
		}
	}
	return false
}

func (cmd *Command) MatchAndExecute(args string, out io.Writer) bool {
	if checkEnvironments(cmd.Environments) && cmd.MatchFn != nil && cmd.MatchFn(cmd.Text, args) {
		cmd.Fn(strings.Split(args, " "), out)
		return true
	}
	return false
}

func (cmd *Command) keyBindFn(buffer *prompt.Buffer) {
	out := &bytes.Buffer{}
	cmd.Fn([]string{}, out)
	buffer.InsertText("\n"+out.String(), false, false)
}

func BuildPromptKeyBinds() []prompt.KeyBind {
	var keybinds []prompt.KeyBind
	for _, cmd := range DefaultCommands {
		x := cmd // for closure
		if cmd.HokKey > 0 {
			keybinds = append(keybinds, prompt.KeyBind{
				Key: x.HokKey,
				Fn:  x.keyBindFn,
			})
		}
	}
	return keybinds
}

func IgnoreCaseMatch(text string, cmd string) bool {
	return strings.EqualFold(cmd, text)
}

func StartWithMatch(text string, cmd string) bool {
	return strings.HasPrefix(cmd, text)
}

func AllMatch(text string, cmd string) bool {
	return true
}

func Logo(args []string, out io.Writer) {
	_, _ = out.Write([]byte("\n ________  ________  ____  _____  _____  ____  ____         ______  _____     _____  \n" +
		"|_   __  ||_   __  ||_   \\|_   _||_   _||_  _||_  _|      .' ___  ||_   _|   |_   _| \n" +
		"  | |_ \\_|  | |_ \\_|  |   \\ | |    | |    \\ \\  / /______ / .'   \\_|  | |       | |   \n" +
		"  |  _|     |  _| _   | |\\ \\| |    | |     > `' <|______|| |         | |   _   | |   \n" +
		" _| |_     _| |__/ | _| |_\\   |_  _| |_  _/ /'`\\ \\_      \\ `.___.'\\ _| |__/ | _| |_  \n" +
		"|_____|   |________||_____|\\____||_____||____||____|      `.____ .'|________||_____| \n" +
		"                                                                                     \n" +
		"                                              https://github.com/fenixsoft/fenix-cli\n" +
		"                                              Press key <F1> to get help information\n" +
		"                                                                                     \n"))
}

func getKeyName(k prompt.Key) string {
	if k >= prompt.F1 && k <= prompt.F12 {
		return "F" + strconv.Itoa(int(k-prompt.F1+1))
	}
	return ""
}

func getEnv(envs []Environment) string {
	if len(envs) != 0 {
		return strings.Join(util.Slice(envs, reflect.TypeOf([]string(nil))).([]string), ",")
	} else {
		return "All"
	}
}

// print help information
func HelpInfo(args []string, out io.Writer) {
	Logo(args, out)
	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"Command", "Hotkey", "Description", "Environment"})
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetColWidth(120)
	table.SetCenterSeparator("|")
	cmds := DefaultCommands
	for _, v := range Environments {
		cmds = append(cmds, v.Commands...)
	}
	for _, v := range cmds {
		table.Append([]string{v.Text, getKeyName(v.HokKey), v.Description, getEnv(v.Environments)})
	}

	// keyboard
	for _, v := range [][]string{
		{"", "", "", ""},
		{"", "Ctrl+A", "Go to the beginning of the line (Home)", ""},
		{"", "Ctrl+E", "Go to the end of the line (End)", ""},
		{"", "Ctrl+P", "Previous command (Up arrow)", ""},
		{"", "Ctrl+N", "Next command (Down arrow)", ""},
		{"", "Ctrl+F", "Forward one character", ""},
		{"", "Ctrl+B", "Backward one character", ""},
		{"", "Ctrl+D", "Delete character under the cursor", ""},
		{"", "Ctrl+H", "Delete character before the cursor (Backspace)", ""},
		{"", "Ctrl+W", "Cut the word before the cursor to the clipboard", ""},
		{"", "Ctrl+K", "Cut the line after the cursor to the clipboard", ""},
		{"", "Ctrl+U", "Cut the line before the cursor to the clipboard", ""},
		{"", "Ctrl+L", "Clear the screen", ""},
	} {
		table.Append(v)
	}
	table.Render()
}

// default commands are valid in all environments
var DefaultCommands []Command

func init() {
	DefaultCommands = []Command{
		{
			Text:        "x-help",
			HokKey:      prompt.F1,
			Description: "Show help information",
			MatchFn:     IgnoreCaseMatch,
			Fn:          HelpInfo,
		},
		{
			Text:        "x-envs",
			HokKey:      prompt.F2,
			Description: "Select an environment",
			MatchFn:     IgnoreCaseMatch,
			Fn: func(args []string, out io.Writer) {
				SelectCurrentEnvironment()
			},
		},
		{
			HokKey:      prompt.F11,
			Description: "Show Fenix-CLI logo and GitHub project",
			Fn:          Logo,
		},
		{
			Text:        "exit",
			HokKey:      prompt.F12,
			Description: "Exit Fenix-CLI command prompt",
			MatchFn:     IgnoreCaseMatch,
			Fn: func(args []string, out io.Writer) {
				_, _ = out.Write([]byte("Bye!\n"))
				os.Exit(0)
			},
		},
	}
}
