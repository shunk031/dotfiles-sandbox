package initialize

import (
	"log"
	"os"

	"github.com/shunk031/dotfiles/cmd/common"
)

func isXCodeInstalled() bool {
	if _, err := os.Stat("/Applications/XCode.app"); !os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

func installXCode() {
	common.PrintInPurple("\n  XCode\n\n")

	if !isXCodeInstalled() {
		cmd := "open macappstores://itunes.apple.com/en/app/xcode/id497799835"
		err := execute(cmd)
		if err != nil {
			log.Fatal(err)
		}
	}
	cmd := "until [ -d \"/Applications/XCode.app\"]; do sleep 5; done"
	msg := "XCode.app"
	err := common.Execute(msg, cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func optOutOfAnalytics() {
	err := execute("brew analytics off")
	common.PrintResult("Homebrew (opt-out of analytics)", err)
	if err != nil {
		log.Fatal(err)
	}
}

func installHomebrew() {
	common.PrintInPurple("\n   Homebrew")

	if !common.CmdExists("brew") {
		cmd := `printf "\n" | /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
		//        └─ simulate the ENTER keypress
		err := execute(cmd)
		common.PrintResult("Homebrew", err)
		if err != nil {
			log.Fatal(err)
		}
	}
	optOutOfAnalytics()
}

func runInitMacOsCommon() error {
	installXCode()
	installHomebrew()
	return nil
}
