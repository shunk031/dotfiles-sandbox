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

func uninstallOldDocker() error {
	// ref. https://docs.docker.com/engine/install/ubuntu/#uninstall-old-versions
	packages := []string{"docker", "docker-engine", "docker.io", "containerd", "runc"}
	for _, p := range packages {
		if common.AptPackageIsInstalled(p) {
			cmd := fmt.Sprintf("sudo apt remove -y %s", p)
			msg := fmt.Sprintf("Remove pakcage %s", p)
			if err := common.Execute(msg, cmd); err != nil {
				return err
			}
		}
	}
	return nil
}

func installDockerUsingRepository() error {

	cmd := `/bin/sh -c "$(curl -fsSL https://get.docker.com)"`
	msg := "Install docker"
	if err := common.Execute(msg, cmd); err != nil {
		return err
	} else {
		return nil
	}
}

func installDocker() error {
	// ref. https://docs.docker.com/engine/install/ubuntu/#install-using-the-convenience-script
	// ref. https://github.com/docker/docker-install
	if err := uninstallOldDocker(); err != nil {
		return err
	}
	if err := installDockerUsingRepository(); err != nil {
		return err
	}
	return nil
}

func installDockerCompose() error {
	cmd := "sudo apt update && sudo apt install -y docker-compose-plugin"
	msg := "Install docker-compose"
	return common.Execute(msg, cmd)
}

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

func InstallDocker() error {
	common.PrintInPurple("\n   Docker\n")
	if err := installDocker(); err != nil {
		return err
	}
	if err := installDockerCompose(); err != nil {
		return err
	}
	if err := installHadolint(); err != nil {
		return err
	}
	return nil
}
