package krew

import (
	"fmt"
	"os"
	"os/exec"
	"sigs.k8s.io/krew/cmd/krew/cmd"
	"strings"
)

func RunAction(args []string) {
	rootCmd := cmd.GetRootCmd()
	rootCmd.SetArgs(args)
	_ = rootCmd.Execute()
}

func IsInstallPlugin(name string) bool {
	ret, _ := cmd.GetInstalledPlugin()
	for _, v := range ret {
		if strings.EqualFold(v.Name, name) {
			return true
		}
	}
	return false

}

func CheckAndInstall(name string) {
	if !IsInstallPlugin(name) {
		fmt.Printf("Fetching plugin: %v \n", name)
		RunAction([]string{"install", name})
	}
}

func GetBinPath() []string {
	path := cmd.GetPath().BinPath() + ":" + os.Getenv("PATH")
	return []string{"PATH=" + path}
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

func IsIstiocltAvailable() bool {
	cmd := exec.Command("istioctl", "version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
