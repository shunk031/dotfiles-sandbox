//go:build darwin

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallTmux() error {
	common.PrintInPurple("\n   Install tmux\n")

	if err := common.BrewInstall("tmux (original)", "tmux", common.BrewOpts{}); err != nil {
		return err
	}
	if err := common.BrewInstall("tmux (pasteboard)", "reattach-to-user-namespace", common.BrewOpts{}); err != nil {
		return err
	}
	if err := common.BrewInstall("cmake", "cmake", common.BrewOpts{}); err != nil {
		return err
	}
	return nil
}
