package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func cloneMecabIpadicNeologd(url string, dir string) error {
	cmd := fmt.Sprintf("git clone --quiet --depth 1 %s %s", url, dir)
	msg := fmt.Sprintf("Clone to %s", dir)
	return common.Execute(msg, cmd)
}

func installMecabIpadicNeologd(dir string) error {
	installCmdPath := filepath.Join(dir, "bin", "install-mecab-ipadic-neologd")
	cmd := fmt.Sprintf("%s -n -y", installCmdPath)
	msg := "Install"
	return common.Execute(msg, cmd)
}

func InstallMecabIpadicNeologd() {
	dir, err := ioutil.TempDir("", "mecab-ipadic-neologd-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	url := "https://github.com/neologd/mecab-ipadic-neologd.git"
	err = cloneMecabIpadicNeologd(url, dir)
	if err != nil {
		log.Fatal(err)
	}

	err = installMecabIpadicNeologd(dir)
	if err != nil {
		log.Fatal(err)
	}
}
