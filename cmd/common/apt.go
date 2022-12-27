package common

import (
	"fmt"
	"os/exec"
)

type AptOpts struct {
	ExtraArgs string
}

func AptPackageIsInstalled(pkg string) bool {
	cmd := exec.Command("dpkg", "-s", pkg)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func AptInstall(pkgReadableName string, pkg string, opts AptOpts) error {
	if !AptPackageIsInstalled(pkg) {
		msg := pkgReadableName
		cmd := fmt.Sprintf("sudo apt-get install -y %s %s", opts.ExtraArgs, pkg)
		return Execute(msg, cmd)

	} else {
		printSuccess(pkgReadableName)
		return nil
	}
}

func AptPurge(pkgReadableName string, pkg string, opts AptOpts) error {
	if AptPackageIsInstalled(pkg) {
		msg := pkgReadableName
		cmd := fmt.Sprintf("sudo apt-get purge -y %s %s", opts.ExtraArgs, pkg)
		return Execute(msg, cmd)
	} else {
		printSuccess(pkgReadableName)
		return nil
	}
}
