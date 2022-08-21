package common

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type BrewOpts struct {
	TapValue string
	Cask     string
	CmdArgs  string
}

func BrewTap(tapValue string) bool {
	cmd := exec.Command("brew", "tap", tapValue)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func BrewList(formula string) bool {
	cmd := exec.Command("brew", "list", formula)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func BrewInstall(formulaReadableName string, formula string, opts BrewOpts) {

	// Check if `Homebrew` is installed.
	if !CmdExists("brew") {
		err := fmt.Errorf("%s ('Homebrew' is not installed.)", formulaReadableName)
		printError(err)
		os.Exit(1)
	}

	// If `brew tap` needs to be executed,
	// check if it executed correctly.
	if opts.TapValue != "" {
		if !BrewTap(opts.TapValue) {
			err := fmt.Errorf("%s ('brew tap %s' failed)", formulaReadableName, opts.TapValue)
			printError(err)
			os.Exit(1)
		}
	}

	// Install the specified formula.
	if BrewList(formula) {
		printSuccess(formulaReadableName)
	} else if opts.Cask != "" {
		msg := formulaReadableName
		cmd := fmt.Sprintf("brew install --cask %s %s", formula, opts.CmdArgs)
		err := Execute(msg, cmd)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		msg := formulaReadableName
		cmd := fmt.Sprintf("brew install %s %s", formula, opts.CmdArgs)
		err := Execute(msg, cmd)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BrewUninstall(formulaReadableName string, formula string) {
	cmd := fmt.Sprintf("brew uninstall %s", formula)
	msg := fmt.Sprintf("Uninstall %s with brew uninstall command", formulaReadableName)
	err := Execute(msg, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
