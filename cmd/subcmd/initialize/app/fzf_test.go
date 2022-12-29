package app

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func cleanupForFzf() {
	if common.CmdExists("fzf") {

		fzfDir := filepath.Join(os.Getenv("HOME"), ".fzf")
		uninstallCmd := filepath.Join(fzfDir, "uninstall")

		err := common.ExecuteCmd(uninstallCmd)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestInstallFzf(t *testing.T) {

	cleanupForFzf()

	InstallFzf()

	expectedFzfDirPath := filepath.Join(os.Getenv("HOME"), ".fzf")
	if _, err := os.Stat(expectedFzfDirPath); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Directory does not found to %s", expectedFzfDirPath)
	}
}
