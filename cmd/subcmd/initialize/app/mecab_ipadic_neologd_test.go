package app

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestInstallMecabIpadicNeologd(t *testing.T) {
	InstallMecabIpadicNeologd()

	expectedDirPath := filepath.Join("/usr", "local", "lib", "mecab", "dic", "mecab-ipadic-neologd")
	if _, err := os.Stat(expectedDirPath); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Directory does not found to %s", expectedDirPath)
	}
}
