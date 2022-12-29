//go:build linux

package app

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func cleanupForGhq() {
	goBinDir := filepath.Join(os.Getenv("GOPATH"), "bin")
	ghqBinPath := filepath.Join(goBinDir, "ghq")

	if _, err := os.Stat(ghqBinPath); err == nil {
		err := os.Remove(ghqBinPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestInstallGhq(t *testing.T) {
	cleanupForGhq()

	InstallGhq()

	if !common.CmdExists("ghq") {
		t.Fatal("Command not found: ghq")
	}

	expectedGhqDirPath := filepath.Join(os.Getenv("HOME"), "ghq")
	if _, err := os.Stat(expectedGhqDirPath); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Directory does not found to %s", expectedGhqDirPath)
	}
}
