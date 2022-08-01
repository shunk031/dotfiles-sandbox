package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

const dotfilesLogo = `
                          /$$                                      /$$
                         | $$                                     | $$
     /$$$$$$$  /$$$$$$  /$$$$$$   /$$   /$$  /$$$$$$      /$$$$$$$| $$$$$$$
    /$$_____/ /$$__  $$|_  $$_/  | $$  | $$ /$$__  $$    /$$_____/| $$__  $$
   |  $$$$$$ | $$$$$$$$  | $$    | $$  | $$| $$  \ $$   |  $$$$$$ | $$  \ $$
    \____  $$| $$_____/  | $$ /$$| $$  | $$| $$  | $$    \____  $$| $$  | $$
    /$$$$$$$/|  $$$$$$$  |  $$$$/|  $$$$$$/| $$$$$$$//$$ /$$$$$$$/| $$  | $$
   |_______/  \_______/   \___/   \______/ | $$____/|__/|_______/ |__/  |__/
                                           | $$
                                           | $$
                                           |__/
                     *** This is dotfiles setup script ***
`

func ShowLogo() {
	fmt.Print(dotfilesLogo)
}

const (
	repository = "shunk031/dotfiles"
	version    = "0.0.1"
)

func main() {
	var (
		// DotfilesOrigin     = "git@github.com:" + repository + ".git"
		DotfilesDirectory  = filepath.Join(os.Getenv("HOME"), ".dotfiles")
		DotfilesTarballURL = filepath.Join("https://github.com", repository, "tarball/master")
		IsSkipQuestions    = true
	)

	ShowLogo()
	common.DownloadDotfiles(DotfilesDirectory, DotfilesTarballURL, IsSkipQuestions)
	RunSubcommand()

}
