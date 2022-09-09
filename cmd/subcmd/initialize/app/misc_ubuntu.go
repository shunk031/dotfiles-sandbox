//go:build linux

package app

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/shunk031/dotfiles/cmd/common"
)

const SHELL = "/bin/bash"

func installHadolint() error {
	dotPath, err := common.GetDotPath()
	if err != nil {
		return err
	}

	out_us, err := exec.Command(SHELL, "-c", "uname -s").Output()
	if err != nil {
		return err
	}
	out_um, err := exec.Command(SHELL, "-c", "uname -m").Output()
	if err != nil {
		return err
	}

	os := strings.TrimSuffix(string(out_us), "\n")
	arch := strings.TrimSuffix(string(out_um), "\n")

	hadolintDir := filepath.Join(dotPath, "bin")
	hadolintURL := fmt.Sprintf("https://github.com/hadolint/hadolint/releases/latest/download/hadolint-%s-%s", os, arch)
	hadolintPath := filepath.Join(hadolintDir, "hadolint")

	cmd := fmt.Sprintf("curl -LsSo %s %s && chmod +x %s", hadolintPath, hadolintURL, hadolintPath)
	msg := fmt.Sprintf("Download %s to install hadolint to %s", hadolintURL, hadolintPath)
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
	common.PrintInPurple("\n   Miscellaneous\n")

	if err := installMisc(); err != nil {
		return err
	}
	if err := installHadolint(); err != nil {
		return err
	}
	return nil
}
