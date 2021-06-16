package krew

import (
	"os/exec"
	"sigs.k8s.io/krew/cmd/krew/cmd"
)

func RunAction(args []string) {
	rootCmd := cmd.GetRootCmd()
	rootCmd.SetArgs(args)
	_ = rootCmd.Execute()
}

func IsTSharkAvailable() bool {
	cmd := exec.Command("tshark", "-v")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func IsXDGAvailable() bool {
	cmd := exec.Command("xdg-open", "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
