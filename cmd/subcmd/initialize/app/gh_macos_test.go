//go:build macos

package app

import (
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func TestInstallGh(t *testing.T) {
	InstallGh()

	if !common.CmdExists("gh") {
		t.Fatal("Package `gh` is not installed")
	}
}
