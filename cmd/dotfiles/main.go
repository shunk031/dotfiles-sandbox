package main

import (
	"fmt"
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
										   
            *** This is the setup script for shunk031's dotfiles ***
`

func ShowLogo() {
	fmt.Print(dotfilesLogo)
}

const (
	repository = "shunk031/dotfiles"
	version    = "0.0.1"
)

func main() {
	// var (
	// 	DotfilesOrigin     = "git@github.com:" + repository + ".git"
	// 	DotfilesDirectory  = filepath.Join(os.Getenv("HOME"), ".dotfiles")
	// 	DotfilesTarballURL = filepath.Join("https://github.com", repository, "tarball/master")
	// 	IsSkipQuestions    = true
	// )

	ShowLogo()
	// common.DownloadDotfiles(DotfilesDirectory, DotfilesTarballURL, IsSkipQuestions)
	RunSubcommand()

}
