//go:build linux

package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func InstallGhq() error {
	common.PrintInPurple("\n   Install ghq\n")

	msg := "ghq"
	cmd := "go install github.com/x-motemen/ghq@latest"
	if err := common.Execute(msg, cmd); err != nil {
		return err
	}

	ghqDir := filepath.Join(os.Getenv("HOME"), "ghq")
	cmd = fmt.Sprintf("mkdir -p %s", ghqDir)
	if err := common.ExecuteCmd(cmd); err != nil {
		return err
	}
	return nil
}
