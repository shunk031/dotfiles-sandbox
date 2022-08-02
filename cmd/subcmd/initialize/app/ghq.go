package app

import (
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func InstallGhq() {
	common.PrintInPurple("\n   Install ghq\n\n")
	common.BrewInstall("ghq", "ghq", "", "", "")
	common.Mkd(filepath.Join(os.Getenv("HOME"), "ghq"))
}
