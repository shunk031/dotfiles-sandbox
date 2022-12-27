//go:build darwin

package app

import (
	"fmt"

	"github.com/shunk031/dotfiles/cmd/common"
)

func runMasInstall(appId string, appName string) error {
	cmd := fmt.Sprintf("mas install %s", appId)
	msg := fmt.Sprintf("Install %s", appName)
	return common.Execute(msg, cmd)
}

func installBandwidthPlus() error {
	appId := "490461369"
	appName := "Bandwidth+"
	return runMasInstall(appId, appName)
}

func InstallMacAppFromStore() error {
	common.PrintInPurple("\n   Install apps from Mac App Store\n")

	// install mas (Mac App Store command line interface)
	common.BrewInstall("mas", "mas", common.BrewOpts{})

	// install apps
	return installBandwidthPlus()
}
