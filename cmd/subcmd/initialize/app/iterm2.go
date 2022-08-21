package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func InstallIterm2() {
	common.PrintInPurple("\n   Install iTerm2\n\n")

	common.BrewInstall("iterm2", "iterm2", common.BrewOpts{TapValue: "homebrew/cask", Cask: "cask"})

	msg := "Open the app so the preference files get initialized."
	cmd := "open -g \"/Applications/iTerm.app\" && sleep 2"
	err := common.Execute(msg, cmd)
	if err != nil {
		log.Fatal(err)
	}

	dotPath, err := common.GetDotPath()
	if err != nil {
		log.Fatal(err)
	}
	jsonFileName := "hotkey_window.json"
	srcJsonPath := filepath.Join(dotPath, "machines", "macos", jsonFileName)
	dstJsonPath := filepath.Join(os.Getenv("HOME"), "Library", "Application\\ Support", "iTerm2", "DynamicProfiles", jsonFileName)

	msg = fmt.Sprintf("Create symbolic link from %s to %s", srcJsonPath, dstJsonPath)
	cmd = fmt.Sprintf("ln -sfn %s %s", srcJsonPath, dstJsonPath)
	fmt.Println(cmd)
	err = common.Execute(msg, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
