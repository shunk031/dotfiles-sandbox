package app

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func TestInstallFzf(t *testing.T) {
	InstallFzf()

	if !common.CmdExists("fzf") {
		t.Fatal("Command not found: fzf")
	}

	expectedFzfDirPath := filepath.Join(os.Getenv("HOME"), ".fzf")
	if _, err := os.Stat(expectedFzfDirPath); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Directory does not found to %s", expectedFzfDirPath)
	}
}
