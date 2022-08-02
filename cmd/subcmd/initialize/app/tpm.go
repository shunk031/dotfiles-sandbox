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
	cmd := fmt.Sprintf("rm -rf %s && git clone --quiet %s %s", dir, url, dir)
	msg := "Install tpm (tmux plugin manager)"
	return common.Execute(msg, cmd)
}

func installTpmPlugins(dir string) error {
	cmd := filepath.Join(dir, "scripts", "install_plugins.sh")
	fmt.Println(cmd)
	msg := "Install tpm plugins"
	return common.Execute(msg, cmd)
}

func installTmuxMemCpuLoad() error {
	dir, err := ioutil.TempDir("", "tmux-cpu-mem-load-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	url := "https://github.com/thewtex/tmux-mem-cpu-load.git"

	cmd := fmt.Sprintf("git clone --quiet %s %s && cd %s && cmake . -DMAKE_INSTALL_PREFIX=%s && make && make install", url, dir, dir, os.Getenv("HOME"))
	msg := "Install tmux mem cpu load"
	return common.Execute(msg, cmd)
}

func InstallTpm() {
	dir := filepath.Join(os.Getenv("HOME"), ".tmux", "plugins", "tpm")
	url := "https://github.com/tmux-plugins/tpm"

	err := installTpm(dir, url)
	if err != nil {
		log.Fatal(err)
	}

	err = installTpmPlugins(dir)
	if err != nil {
		log.Fatal(err)
	}

	err = installTmuxMemCpuLoad()
	if err != nil {
		log.Fatal(err)
	}

}
