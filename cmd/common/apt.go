package common

import (
	"fmt"
	"os/exec"
)

type AptOpts struct {
	ExtraArgs string
}

func packageIsInstalled(pkg string) bool {
	cmd := exec.Command("dpkg", "-s", pkg)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func AptInstall(pkgReadableName string, pkg string, opts AptOpts) error {
	if !packageIsInstalled(pkg) {
		msg := pkgReadableName
		cmd := fmt.Sprintf("sudo apt-get install -y %s %s", opts.ExtraArgs, pkg)
		return Execute(msg, cmd)

	} else {
		printSuccess(pkgReadableName)
		return nil
	}
}
