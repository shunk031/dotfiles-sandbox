//go:build darwin

package app

import (
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func TestInstallTmux(t *testing.T) {
	InstallTmux()

	if !common.CmdExists("tmux") {
		t.Fatal("Command not found: tmux")
	}

	if !common.CmdExists("reattach-to-user-namespace") {
		t.Fatal("Command not found: reattach-to-user-namespace")
	}

	if !common.CmdExists("cmake") {
		t.Fatal("Command not found: cmake")
	}
}
