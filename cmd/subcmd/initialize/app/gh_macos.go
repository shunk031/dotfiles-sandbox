//go:build macos

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallGh() {
	common.PrintInPurple("\n   Install gh\n")

	common.BrewInstall("gh command", "gh", common.BrewOpts{})
}
