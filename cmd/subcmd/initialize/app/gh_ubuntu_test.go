//go:build linux

package app

import (
	"log"
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func cleanupForGh() {
	if err := common.AptPurge("gh", "gh", common.AptOpts{}); err != nil {
		log.Fatal(err)
	}
}

func TestInstallGh(t *testing.T) {
	cleanupForGh()

	InstallGh()

	if !common.CmdExists("gh") {
		t.Fatal("Package `gh` is not installed")
	}
}
