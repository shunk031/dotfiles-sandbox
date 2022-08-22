//go:build amd64

package initialize

import (
	"github.com/shunk031/dotfiles/cmd/subcmd/initialize/app"
)

func runInitMacOsArch() error {

	app.InstallIterm2()
	app.InstallTmux()

	app.InstallMecab()
	app.InstallMecabIpadicNeologd()

	app.InstallPowerlevel10kRequirements()

	return nil
}
