//go:build darwin

package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func InstallGhq() error {
	common.PrintInPurple("\n   Install ghq\n")

	if err := common.BrewInstall("ghq", "ghq", common.BrewOpts{}); err != nil {
		return err
	}

	ghqDir := filepath.Join(os.Getenv("HOME"), "ghq")
	cmd := fmt.Sprintf("mkdir -p %s", ghqDir)
	if err := common.ExecuteCmd(cmd); err != nil {
		return err
	}
	return nil
}
