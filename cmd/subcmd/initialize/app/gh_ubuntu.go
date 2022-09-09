//go:build linux

package app

import (
	"strings"

	"github.com/shunk031/dotfiles/cmd/common"
)

func installGh() error {

	cmds := []string{
		"curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg",
		"sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg",
		"echo \"deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main\" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null",
		"sudo apt update",
		"sudo apt install gh",
	}
	cmd := strings.Join(cmds, "&&")

	msg := "Install gh"
	if err := common.Execute(msg, cmd); err != nil {
		return err
	} else {
		return nil
	}
}

func InstallGh() error {
	common.PrintInPurple("\n   Install gh\n")
	if err := installGh(); err != nil {
		return err
	}
	return nil
}
