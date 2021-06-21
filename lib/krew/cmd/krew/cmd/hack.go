package cmd

import (
	"sigs.k8s.io/krew/internal/environment"
	"sigs.k8s.io/krew/internal/installation"
	"sigs.k8s.io/krew/pkg/index"
)

func GetInstalledPlugin() ([]index.Receipt, error) {
	return installation.GetInstalledPluginReceipts(paths.InstallReceiptsPath())
}

func GetPath() environment.Paths {
	return environment.MustGetKrewPaths()
}
