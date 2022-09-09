//go:build darwin

package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func installNerdFont() error {
	fontDir := filepath.Join(os.Getenv("HOME"), "Library", "Fonts")
	fontURL := "https://github.com/ryanoasis/nerd-fonts/raw/master/patched-fonts/RobotoMono/Medium/complete/Roboto%20Mono%20Medium%20Nerd%20Font%20Complete%20Mono.ttf"
	fontName := "Roboto\\ Mono\\ Nerd\\ Font\\ Complete.ttf"

	// install curl dependency
	common.BrewInstall("curl", "curl", common.BrewOpts{})

	cmd := fmt.Sprintf("mkdir -p %s && cd %s && curl -fLo %s %s", fontDir, fontDir, fontName, fontURL)
	msg := fmt.Sprintf("Install %s to %s", fontName, fontDir)

	if err := common.Execute(msg, cmd); err != nil {
		return err
	} else {
		return nil
	}
}

func InstallPowerlevel10kRequirements() error {
	common.PrintInPurple("\n   Install powerlevel10k requirements\n\n")

	if err := installNerdFont(); err != nil {
		return err
	} else {
		return nil
	}
}
