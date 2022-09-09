//go:build darwin

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallMecab() error {
	common.PrintInPurple("\n   Install mecab\n")

	if err := common.BrewInstall("mecab", "mecab", common.BrewOpts{}); err != nil {
		return err
	}
	if err := common.BrewInstall("mecab-ipadic", "mecab-ipadic", common.BrewOpts{}); err != nil {
		return err
	}

	return nil
}
