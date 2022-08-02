package common

import (
	"log"
	"os"
	"testing"
)

func TestGetDotPath(t *testing.T) {
	os.Setenv("DOTPATH", "")
	dotPath, err := GetDotPath()
	if err == nil {
		t.Fatal(err)
	}
	if dotPath != "" {
		t.Fatal("Test failed")
	}

	os.Setenv("DOTPATH", "/path/to/dotfiles")
	dotPath, err = GetDotPath()
	if err != nil {
		log.Fatal("Test failed")
	}
	if dotPath == "" {
		t.Fatal("Test failed")
	}
}

func TestCmdExists(t *testing.T) {
	if !CmdExists("ls") {
		t.Fatal("Test failed")
	}

	if CmdExists("command-probably-never-exists") {
		t.Fatal("Test failed")
	}
}
