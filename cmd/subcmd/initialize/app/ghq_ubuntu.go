//go:build linux

package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func installGhq() error {
	msg := "ghq command"
	cmd := "go install github.com/x-motemen/ghq@latest"
	return common.Execute(msg, cmd)
}

func mkdGhq() error {
	ghqDir := filepath.Join(os.Getenv("HOME"), "ghq")
	cmd := fmt.Sprintf("mkdir -p %s", ghqDir)
	return common.ExecuteCmd(cmd)
}

func InstallGhq() error {
	common.PrintInPurple("\n   Install ghq\n\n")

	if err := installGhq(); err != nil {
		return err
	}
	return mkdGhq()
}
