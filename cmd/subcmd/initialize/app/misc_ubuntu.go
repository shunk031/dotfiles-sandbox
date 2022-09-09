//go:build linux

package app

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func installHadolint() error {
	dotPath, err := common.GetDotPath()
	if err != nil {
		return err
	}

	out_us, err := exec.Command("uname -s").Output()
	if err != nil {
		return nil
	}
	out_um, err := exec.Command("uname -m").Output()
	if err != nil {
		return nil
	}

	hadolintDir := filepath.Join(dotPath, "bin")
	hadolintURL := fmt.Sprintf("https://github.com/hadolint/hadolint/releases/latest/download/hadolint-%s-%s", string(out_us), string(out_um))
	hadolintPath := filepath.Join(hadolintDir, "hadolint")

	cmd := fmt.Sprintf("curl -osL %s %s && chmod +x %s", hadolintPath, hadolintURL, hadolintPath)
	msg := fmt.Sprintf("Install hadolint to %s", hadolintPath)
	if err := common.Execute(msg, cmd); err != nil {
		return err
	} else {
		return nil
	}
}

func installMisc() error {
	packages := []string{"zsh", "shellcheck", "exa"}
	for _, p := range packages {
		if err := common.AptInstall(p, p, common.AptOpts{}); err != nil {
			return err
		}
	}
	return nil
}

func InstallMisc() error {
	if err := installMisc(); err != nil {
		return err
	}
	if err := installHadolint(); err != nil {
		return err
	}
	return nil
}
