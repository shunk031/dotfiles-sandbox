//go:build darwin

package app

import (
	"log"
	"os/exec"

	"github.com/shunk031/dotfiles/cmd/common"
)

func installGh() {
	common.BrewInstall("gh command", "gh", common.BrewOpts{})
}

func runGhAuthLogin() {
	cmd := "gh auth login -h github.com -p https"

	c := exec.Command(cmd)
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}

	// err := common.ExecuteCmd(cmd)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func InstallGh() {
	common.PrintInPurple("\n   Install gh\n")

	installGh()
	// runGhAuthLogin()
}
