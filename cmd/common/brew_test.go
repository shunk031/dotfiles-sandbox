package common

import "testing"

func TestBrewTap(t *testing.T) {
	if !BrewTap("homebrew/cask") {
		t.Fatal("Cannot find `homebrew/cask` with brew tap command")
	}

	if BrewTap("tap-probably-never-exists") {
		t.Fatal("Test failed")
	}
}

func TestBrewList(t *testing.T) {
	if !BrewList("gcc") {
		t.Fatal("Cannot find `gcc` with brew list command")
	}
	if BrewList("package-probably-never-exists") {
		t.Fatal("Test failed")
	}
}

func TestBrewInstall(t *testing.T) {
	BrewInstall("formula-probably-never-exists", "formula-probably-never-exists", &BrewOpts{})
}
