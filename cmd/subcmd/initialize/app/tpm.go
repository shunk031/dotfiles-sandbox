package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func installTpm(dir string, url string) error {
	cmd := fmt.Sprintf("rm -rf %s && git clone %s %s", dir, url, dir)
	msg := "Install tpm (tmux plugin manager)"
	return common.Execute(msg, cmd)
}

func installTpmPlugins(dir string) error {
	cmd := filepath.Join(dir, "scripts", "install_plugins.sh")
	msg := "Install tpm plugins"
	return common.Execute(msg, cmd)
}

func installTmuxMemCpuLoad() error {
	dir, err := ioutil.TempDir("", "tmux-mem-cpu-load-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	url := "https://github.com/thewtex/tmux-mem-cpu-load.git"

	cmd := fmt.Sprintf("git clone %s %s && cd %s && cmake . && make && sudo make install", url, dir, dir)
	msg := "Install tmux mem cpu load"
	return common.Execute(msg, cmd)
}

func InstallTpm() error {
	dir := filepath.Join(os.Getenv("HOME"), ".tmux", "plugins", "tpm")
	url := "https://github.com/tmux-plugins/tpm"

	if err := installTpm(dir, url); err != nil {
		return err
	}
	if err := installTpmPlugins(dir); err != nil {
		return err
	}
	if err := installTmuxMemCpuLoad(); err != nil {
		return err
	}
	return nil
}
