package common

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func brewTap(tapValue string) bool {
	cmd := exec.Command("brew", "tap", tapValue)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func brewList(formula string) bool {
	cmd := exec.Command("brew", "list", formula)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func BrewInstall(formulaReadableName string, formula string, tapValue string, cask string, cmdArgs string) {

	// Check if `Homebrew` is installed.
	if !CmdExists("brew") {
		err := fmt.Errorf("%s ('Homebrew' is not installed.)", formulaReadableName)
		printError(err)
		os.Exit(1)
	}

	// If `brew tap` needs to be executed,
	// check if it executed correctly.
	if tapValue != "" {
		if !brewTap(tapValue) {
			err := fmt.Errorf("%s ('brew tap %s' failed)", formulaReadableName, tapValue)
			printError(err)
			os.Exit(1)
		}
	}

	// Install the specified formula.
	if brewList(formula) {
		printSuccess(formulaReadableName)
	} else if cask != "" {
		msg := formulaReadableName
		cmd := fmt.Sprintf("brew install --cask %s %s", formula, cmdArgs)
		err := Execute(msg, cmd)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		msg := formulaReadableName
		cmd := fmt.Sprintf("brew install %s %s", formula, cmdArgs)
		err := Execute(msg, cmd)
		if err != nil {
			log.Fatal(err)
		}
	}
}
