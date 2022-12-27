//go:build linux

package app

import (
	"log"
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func cleanupForTmux() {
	if err := common.AptPurge("tmux (original)", "tmux", common.AptOpts{}); err != nil {
		log.Fatal(err)
	}

	if err := common.AptPurge("tmux (pastebord)", "xsel", common.AptOpts{}); err != nil {
		log.Fatal(err)
	}

	if err := common.AptPurge("cmake", "cmake", common.AptOpts{}); err != nil {
		log.Fatal(err)
	}
}

func TestInstallTmux(t *testing.T) {

	cleanupForTmux()

	err := InstallTmux()
	if err != nil {
		t.Fatal(err)
	}

	if !common.CmdExists("tmux") {
		t.Fatal("Command not found: tmux")
	}

	if !common.CmdExists("xsel") {
		t.Fatal("Command not found: xsel")
	}

	if !common.CmdExists("cmake") {
		t.Fatal("Command not found: cmake")
	}
}
