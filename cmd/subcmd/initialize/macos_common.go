//go:build darwin

package initialize

import (
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

func installXCode() error {
	common.PrintInPurple("\n  XCode\n\n")

	if !isXCodeInstalled() {
		cmd := "open macappstores://itunes.apple.com/en/app/xcode/id497799835"
		if err := common.ExecuteCmd(cmd); err != nil {
			return err
		}
	}
	cmd := "until [ -d \"/Applications/XCode.app\"]; do sleep 5; done"
	msg := "XCode.app"
	if err := common.Execute(msg, cmd); err != nil {
		return err
	} else {
		return nil
	}
}

func optOutOfAnalytics() error {
	err := common.ExecuteCmd("brew analytics off")
	common.PrintResult("Homebrew (opt-out of analytics)", err)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func installHomebrew() error {
	common.PrintInPurple("\n   Homebrew")

	if !common.CmdExists("brew") {
		cmd := `printf "\n" | /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
		//        └─ simulate the ENTER keypress
		err := common.ExecuteCmd(cmd)
		common.PrintResult("Homebrew", err)
		if err != nil {
			return err
		}
	}
	if err := optOutOfAnalytics(); err != nil {
		return err
	} else {
		return nil
	}
}

func runInitMacOsCommon() error {
	installXCode()
	installHomebrew()
	return nil
}
