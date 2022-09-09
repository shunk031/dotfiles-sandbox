//go:build linux

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallMecab() error {
	common.PrintInPurple("\n   Install mecab\n")

	if err := common.AptInstall("apt install mecab", "mecab", common.AptOpts{}); err != nil {
		return err
	}
	if err := common.AptInstall("apt install libmecab-dev", "libmecab-dev", common.AptOpts{}); err != nil {
		return err
	}
	if err := common.AptInstall("apt install mecab-ipadict-uft8", "mecab-ipadic-utf8", common.AptOpts{}); err != nil {
		return err
	}
	return nil
}
