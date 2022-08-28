//go:build macos

package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func InstallGhq() {
	common.PrintInPurple("\n   Install ghq\n")
	common.BrewInstall("ghq", "ghq", common.BrewOpts{})

	ghqDir := filepath.Join(os.Getenv("HOME"), "ghq")
	cmd := fmt.Sprintf("mkdir -p %s", ghqDir)
	common.ExecuteCmd(cmd)
}
