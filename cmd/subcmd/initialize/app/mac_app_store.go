//go:build darwin

package app

import (
	"fmt"

	"github.com/shunk031/dotfiles/cmd/common"
)

func installMas() error {
	return common.BrewInstall("mas", "mas", common.BrewOpts{})
}

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
	if err := installMas(); err != nil {
		return err
	}

	// install apps
	return installBandwidthPlus()
}
