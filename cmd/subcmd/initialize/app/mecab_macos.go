//go:build macos

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallMecab() {
	common.PrintInPurple("\n   Install mecab\n\n")

	// install mecab
	common.BrewInstall("mecab", "mecab", "", "", "")

	// install mecab-ipadic
	common.BrewInstall("mecab-ipadic", "mecab-ipadic", "", "", "")
}
