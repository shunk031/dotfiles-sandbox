package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shunk031/dotfiles/cmd/subcmd/clean"
	"github.com/shunk031/dotfiles/cmd/subcmd/deploy"
	"github.com/shunk031/dotfiles/cmd/subcmd/initialize"
	"github.com/shunk031/dotfiles/cmd/subcmd/install"
	"github.com/shunk031/dotfiles/cmd/subcmd/local"
	"github.com/shunk031/dotfiles/cmd/subcmd/update"
)

func RunSubcommand() {

	if len(os.Args) < 2 {
		fmt.Println("Expected `init`, `deploy`, `update`, `local`, `clean`, `install` subcommands")
		os.Exit(1)
	}

	var err error
	switch os.Args[1] {
	case "init":
		err = initialize.RunInitCmd()
	case "deploy":
		err = deploy.RunDeployCmd()
	case "update":
		err = update.RunUpdateCmd()
	case "local":
		err = local.RunLocalCmd()
	case "clean":
		err = clean.RunCleanCmd()
	case "install":
		err = install.RunInstallCmd()
	default:
		fmt.Println("Expected `init`, `deploy`, `update`, `local`, `clean`, `install` subcommands")
		os.Exit(1)
	}
	if err != nil {
		log.Fatal(err)
	}
}
