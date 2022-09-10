//go:build linux

package app

import (
	"github.com/shunk031/dotfiles/cmd/common"
)

func installMisc() error {
	packages := []string{"zsh", "shellcheck", "exa"}
	for _, p := range packages {
		if err := common.AptInstall(p, p, common.AptOpts{}); err != nil {
			return err
		}
	}
	return nil
}

func InstallMisc() error {
	common.PrintInPurple("\n   Miscellaneous\n")

	if err := installMisc(); err != nil {
		return err
	}
	return nil
}
