//go:build macos

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallTmux() {
	common.PrintInPurple("\n   tmux\n\n")
	common.BrewInstall("tmux (original)", "tmux", "", "", "")
	common.BrewInstall("tmux (pasteboard)", "reattach-to-user-namespace", "", "", "")
	common.BrewInstall("cmake", "cmake", "", "", "")
}
