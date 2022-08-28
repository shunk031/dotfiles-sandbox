//go:build ubuntu

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallTmux() {
	common.PrintInPurple("\n   Install tmux\n")

	common.AptInstall("tmux (original)", "tmux", common.AptOpts{})
	common.AptInstall("tmux (pastebord)", "xsel", common.AptOpts{})
	common.AptInstall("cmake", "cmake", common.AptOpts{})
}
