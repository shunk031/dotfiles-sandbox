//go:build linux

package app

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestInstallPowerlevel10kRequirements(t *testing.T) {
	InstallPowerlevel10kRequirements()

	downloadedFontName := "Roboto Mono Nerd Font Complete.ttf"
	downloadedFontPath := filepath.Join(os.Getenv("HOME"), ".local", "share", "fonts", downloadedFontName)
	if _, err := os.Stat(downloadedFontPath); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Font file %s does not exists.", downloadedFontName)
	}
}
