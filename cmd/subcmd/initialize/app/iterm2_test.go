//go:build macos

package app

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestInstallIterm2(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	basepath := filepath.Join(filepath.Dir(dir), "..", "..", "..")
	err = os.Setenv("DOTPATH", basepath)
	if err != nil {
		log.Fatal(err)
	}
	InstallIterm2()

	itermAppPath := filepath.Join("/Applications", "iTerm.app")
	if _, err := os.Stat(itermAppPath); errors.Is(err, os.ErrNotExist) {
		t.Fatal("App not found: iTerm.app")
	}

	profileJsonName := "hotkey_window.json"
	profileJsonPath := filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "iTerm2", "DynamicProfiles", profileJsonName)
	if _, err := os.Stat(profileJsonPath); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Json file for dynamic profiles is not found in %s", profileJsonPath)
	}
}
