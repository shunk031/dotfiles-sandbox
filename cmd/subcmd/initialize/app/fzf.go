package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func cloneFzf(dir string, url string) error {
	cmd := fmt.Sprintf("rm -rf %s && git clone %s %s", dir, url, dir)
	msg := fmt.Sprintf("Clone to %s", dir)
	return common.Execute(msg, cmd)
}

func installFzf(dir string) error {
	installCmdPath := filepath.Join(dir, "install")
	cmd := fmt.Sprintf("%s --key-bindings --completion --no-update-rc", installCmdPath)
	msg := fmt.Sprintf("Finish: %s", cmd)
	return common.Execute(msg, cmd)
}

func InstallFzf() error {
	common.PrintInPurple("\n   Install fzf\n")

	fzfDir := filepath.Join(os.Getenv("HOME"), ".fzf")
	fzfUrl := "https://github.com/junegunn/fzf.git"
	if err := cloneFzf(fzfDir, fzfUrl); err != nil {
		return err
	}
	if err := installFzf(fzfDir); err != nil {
		return err
	}
	return nil
}
