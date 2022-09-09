//go:build linux

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallMecabIpadicNeologdRequirements() error {
	requirements := []string{"git", "make", "curl", "xz-utils", "file"}
	for _, req := range requirements {
		if err := common.AptInstall(req, req, common.AptOpts{}); err != nil {
			return err
		}
	}
	return nil
}
