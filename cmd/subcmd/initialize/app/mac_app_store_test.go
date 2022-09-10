//go:build darwin

package app

import (
	"path/filepath"
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func TestInstallMas(t *testing.T) {
	if err := installMas(); err != nil {
		t.Fatal(err)
	}
	if !common.CmdExists("mas") {
		t.Fatal("Command not found: mas")
	}
}

func TestInstallBandwidthPlus(t *testing.T) {
	if err := installBandwidthPlus(); err != nil {
		t.Fatal(err)
	}
	appPath := filepath.Join("/Applications", "Bandwidth+.app")
	if !common.PathExists(appPath) {
		t.Fatal("App not found: Bandwidth+.app")
	}
}
