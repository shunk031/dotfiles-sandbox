//go:build linux

package initialize

import "github.com/shunk031/dotfiles/cmd/subcmd/initialize/app"

func runInitUbuntuCommon() error {
	if err := app.InstallMisc(); err != nil {
		return err
	}
	if err := app.InstallTmux(); err != nil {
		return err
	}
	if err := app.InstallGhq(); err != nil {
		return err
	}
	if err := app.InstallMecab(); err != nil {
		return err
	}
	if err := app.InstallMecabIpadicNeologd(); err != nil {
		return err
	}
	if err := app.InstallPowerlevel10kRequirements(); err != nil {
		return err
	}
	return nil
}
