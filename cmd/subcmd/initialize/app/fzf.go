package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func cloneFzf(dir string, url string) error {
	cmd := fmt.Sprintf("rm -rf %s && git clone --quiet %s %s", dir, url, dir)
	msg := fmt.Sprintf("Clone to %s", dir)
	return common.Execute(msg, cmd)
}

func installFzf(dir string) error {
	installCmdPath := filepath.Join(dir, "install")
	cmd := fmt.Sprintf("%s --key-bindings --completion --no-update-rc", installCmdPath)
	msg := cmd
	return common.Execute(msg, cmd)
}

func InstallFzf() {

	fzfDir := filepath.Join(os.Getenv("HOME"), ".fzf")
	fzfUrl := "https://github.com/junegunn/fzf.git"
	err := cloneFzf(fzfDir, fzfUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = installFzf(fzfDir)
	if err != nil {
		log.Fatal(err)
	}
}
