//go:build darwin

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallTmux() {
	common.PrintInPurple("\n   Install tmux\n")
	common.BrewInstall("tmux (original)", "tmux", common.BrewOpts{})
	common.BrewInstall("tmux (pasteboard)", "reattach-to-user-namespace", common.BrewOpts{})
	common.BrewInstall("cmake", "cmake", common.BrewOpts{})
}
