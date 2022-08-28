//go:build ubuntu

package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func installNerdFont() {
	fontDir := filepath.Join(os.Getenv("HOME"), ".local", "share", "fonts")
	fontURL := "https://github.com/ryanoasis/nerd-fonts/raw/master/patched-fonts/RobotoMono/Medium/complete/Roboto%20Mono%20Medium%20Nerd%20Font%20Complete%20Mono.ttf"
	fontName := "Roboto\\ Mono\\ Nerd\\ Font\\ Complete.ttf"

	// install curl dependency
	common.AptInstall("curl", "curl", common.AptOpts{})

	cmd := fmt.Sprintf("mkdir -p %s && cd %s && curl -fLo %s %s", fontDir, fontDir, fontName, fontURL)
	msg := fmt.Sprintf("Install %s to %s", fontName, fontDir)

	err := common.Execute(msg, cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func InstallPowerlevel10kRequirements() {
	common.PrintInPurple("\n   Install powerlevel10k requirements\n")
	installNerdFont()
}
