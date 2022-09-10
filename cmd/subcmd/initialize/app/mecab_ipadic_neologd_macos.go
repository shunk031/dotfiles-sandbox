//go:build darwin

package app

import "github.com/shunk031/dotfiles/cmd/common"

func installMecabIpadicNeologdRequirements() error {
	requirements := []string{"git", "curl", "xz"}
	for _, req := range requirements {
		if err := common.BrewInstall(req, req, common.BrewOpts{}); err != nil {
			return err
		}
	}
	return nil
}
