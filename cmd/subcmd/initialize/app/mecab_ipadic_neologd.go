package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func isMecabIpadicNeologdExists() bool {
	cmd := exec.Command("/bin/bash", "-c", "echo `mecab-config --dicdir`/mecab-ipadic-neologd")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	neologdPath := string(out)
	fmt.Println(neologdPath)

	if _, err := os.Stat(neologdPath); os.IsNotExist(err) {
		// fmt.Println(f.IsDir())
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func cloneMecabIpadicNeologd(url string, dir string) error {
	cmd := fmt.Sprintf("git clone --depth 1 %s %s", url, dir)
	msg := fmt.Sprintf("Clone to %s", dir)
	return common.Execute(msg, cmd)
}

func installMecabIpadicNeologd(dir string) error {
	installCmdPath := filepath.Join(dir, "bin", "install-mecab-ipadic-neologd")
	cmd := fmt.Sprintf("%s -n -y", installCmdPath)
	msg := "Install"
	return common.Execute(msg, cmd)
}

func InstallMecabIpadicNeologd() error {
	common.PrintInPurple("\n   Install mecab-ipadic-neologd\n")

	dir, err := ioutil.TempDir("", "mecab-ipadic-neologd-")
	if err != nil {
		return err
	}
	// defer os.RemoveAll(dir)

	url := "https://github.com/neologd/mecab-ipadic-neologd.git"
	if err := cloneMecabIpadicNeologd(url, dir); err != nil {
		return err
	}
	if err := installMecabIpadicNeologd(dir); err != nil {
		return err
	}
	return nil
}
