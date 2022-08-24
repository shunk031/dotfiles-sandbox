package app

import (
	"fmt"
	"log"

	"github.com/shunk031/dotfiles/cmd/common"
)

func runMasInstall(appId string, appName string) {
	cmd := fmt.Sprintf("mas install %s", appId)
	msg := fmt.Sprintf("Install %s", appName)
	err := common.Execute(msg, cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func installBandwidthPlus() {
	appId := "490461369"
	appName := "Bandwidth+"
	runMasInstall(appId, appName)
}

func InstallMacAppFromStore() {
	common.PrintInPurple("\n   Install apps from Mac App Store\n")

	// install mas (Mac App Store command line interface)
	common.BrewInstall("mas", "mas", common.BrewOpts{})

	// install apps
	installBandwidthPlus()

}
