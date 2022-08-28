//go:build ubuntu

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallMecab() {
	common.PrintInPurple("\n   Install mecab\n")
	common.AptInstall("mecab", "mecab", common.AptOpts{})
	common.AptInstall("libmecab-dev", "libmecab-dev", common.AptOpts{})
	common.AptInstall("mecab-ipadict-uft8", "mecab-ipadic-utf8", common.AptOpts{})
}
