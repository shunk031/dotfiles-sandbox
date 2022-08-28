//go:build ubuntu

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallGhq() {
	common.PrintInPurple("\n   Install ghq\n\n")

	msg := "ghq"
	cmd := "go install github.com/x-motemen/ghq@latest"
	common.Execute(msg, cmd)
}
