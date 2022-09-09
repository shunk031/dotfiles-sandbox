//go:build linux

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallTmux() error {
	common.PrintInPurple("\n   Install tmux\n")

	if err := common.AptInstall("tmux (original)", "tmux", common.AptOpts{}); err != nil {
		return err
	}
	if err := common.AptInstall("tmux (pastebord)", "xsel", common.AptOpts{}); err != nil {
		return err
	}
	if err := common.AptInstall("cmake", "cmake", common.AptOpts{}); err != nil {
		return err
	}
	if err := InstallTpm(); err != nil {
		return err
	}
	return nil
}
