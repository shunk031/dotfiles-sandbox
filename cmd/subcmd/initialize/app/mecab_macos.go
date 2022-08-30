//go:build darwin

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallMecab() {
	common.PrintInPurple("\n   Install mecab\n")

	// install mecab
	common.BrewInstall("mecab", "mecab", common.BrewOpts{})

	// install mecab-ipadic
	common.BrewInstall("mecab-ipadic", "mecab-ipadic", common.BrewOpts{})
}
