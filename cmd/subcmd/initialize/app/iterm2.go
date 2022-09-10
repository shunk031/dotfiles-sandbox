//go:build darwin

package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

const (
	iterm2ConfigName = "hotkey_window.json"
)

func InstallIterm2() error {
	common.PrintInPurple("\n   Install iTerm2\n")

	common.BrewInstall("iterm2", "iterm2", common.BrewOpts{TapValue: "homebrew/cask", Cask: "cask"})

	msg := "Open the app so the preference files get initialized."
	cmd := "open -g \"/Applications/iTerm.app\" && sleep 2"
	err := common.Execute(msg, cmd)
	if err != nil {
		return err
	}

	dotPath, err := common.GetDotPath()
	if err != nil {
		return err
	}
	srcJsonPath := filepath.Join(dotPath, "machines", "macos", iterm2ConfigName)
	dstJsonPath := filepath.Join(os.Getenv("HOME"), "Library", "Application\\ Support", "iTerm2", "DynamicProfiles", iterm2ConfigName)

	msg = fmt.Sprintf("Create symbolic link from %s to %s", srcJsonPath, dstJsonPath)
	cmd = fmt.Sprintf("ln -sfn %s %s", srcJsonPath, dstJsonPath)

	err = common.Execute(msg, cmd)
	if err != nil {
		return err
	}
	return nil
}
