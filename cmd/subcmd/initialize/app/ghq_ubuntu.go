//go:build linux

package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func InstallGhq() error {
	common.PrintInPurple("\n   Install ghq\n\n")

	err := common.Execute("ghq", "go install github.com/x-motemen/ghq@latest")
	if err != nil {
		return err
	}

	ghqDir := filepath.Join(os.Getenv("HOME"), "ghq")
	cmd := fmt.Sprintf("mkdir -p %s", ghqDir)
	err = common.ExecuteCmd(cmd)
	if err != nil {
		return err
	} else {
		return nil
	}
}
