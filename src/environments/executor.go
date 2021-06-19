package environments

import (
	"bytes"
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"os/exec"
	"strings"
)

func GetDefaultExecutor(runtime string, callback func(), env ...string) prompt.Executor {
	return func(s string) {
		Executor(runtime, s, env...)
		if callback != nil {
			callback()
		}
	}
}

func Executor(program string, args string, env ...string) {
	args = strings.TrimSpace(args)
	if args == "" {
		return
	}

	cmds := DefaultCommands
	for _, env := range Environments {
		cmds = append(cmds, env.Commands...)
	}
	for _, cmd := range cmds {
		if cmd.MatchAndExecute(args, os.Stdout) {
			return
		}
	}

	cmd := exec.Command("/bin/sh", "-c", program+" "+args)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, env...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}
	return
}

func ExecuteAndGetResult(program string, args string, env ...string) (int, string) {
	args = strings.TrimSpace(args)
	if args == "" {
		return -1, ""
	}

	out := &bytes.Buffer{}
	cmd := exec.Command("/bin/sh", "-c", program+" "+args)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, env...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode(), string(out.Bytes())
		}
		return -1, err.Error()
	}
	return 0, string(out.Bytes())
}
