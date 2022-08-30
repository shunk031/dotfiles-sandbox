//go:build darwin

package app

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func TestInstallGhq(t *testing.T) {
	InstallGhq()

	if !common.CmdExists("ghq") {
		t.Fatal("Command not found: ghq")
	}

	expectedGhqDirPath := filepath.Join(os.Getenv("HOME"), "ghq")
	if _, err := os.Stat(expectedGhqDirPath); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Directory does not found to %s", expectedGhqDirPath)
	}
}
