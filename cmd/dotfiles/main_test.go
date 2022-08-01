package main

import (
	"os"
	"testing"
)

func TestPackageIsInstalled(t *testing.T) {
	if !packageIsInstalled("git") {
		t.Fatal("Test failed")
	}

	if packageIsInstalled("hogehoge") {
		t.Fatal("Test failed")
	}
}

func TestIsDebug(t *testing.T) {
	os.Setenv("DEBUG_DOTFILES", "")
	if isDebug() {
		t.Fatal("Test failed")
	}

	os.Setenv("DEBUG_DOTFILES", "1")
	if !isDebug() {
		t.Fatal("Test failed")
	}
}
