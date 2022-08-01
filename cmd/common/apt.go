package common

import (
	"fmt"
	"log"
	"os/exec"
)

func packageIsInstalled(pkg string) bool {
	cmd := exec.Command("dpkg", "-s", pkg)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func AptInstall(pkgReadableName string, pkg string, extraArgs string) {
	if !packageIsInstalled(pkg) {
		msg := pkgReadableName
		cmd := fmt.Sprintf("sudo apt-get install --allow-unauthenticated -qqy %s %s", extraArgs, pkg)
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
