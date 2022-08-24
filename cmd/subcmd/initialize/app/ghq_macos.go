//go:build macos

package app

import (
	"github.com/shunk031/dotfiles/cmd/common"
)

func InstallGhq() {
	common.PrintInPurple("\n   Install ghq\n\n")
	common.BrewInstall("ghq", "ghq", common.BrewOpts{})
	// common.Mkd(filepath.Join(os.Getenv("HOME"), "ghq"))
}
