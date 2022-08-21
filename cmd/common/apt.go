package common

import (
	"fmt"
	"log"
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

func AptInstall(pkgReadableName string, pkg string, opts AptOpts) {
	if !packageIsInstalled(pkg) {
		msg := pkgReadableName
		cmd := fmt.Sprintf("sudo apt-get install --allow-unauthenticated -qqy %s %s", opts.ExtraArgs, pkg)
		//                                                suppress output ─┘│
		//                      assume "yes" as the answer to all prompts ──┘
		err := Execute(msg, cmd)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		printSuccess(pkgReadableName)
	}
}
