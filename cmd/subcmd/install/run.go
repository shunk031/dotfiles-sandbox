package install

import (
	"log"

	"github.com/shunk031/dotfiles/cmd/subcmd/deploy"
	"github.com/shunk031/dotfiles/cmd/subcmd/initialize"
	"github.com/shunk031/dotfiles/cmd/subcmd/local"
)

func RunInstallCmd() error {
	// err := runDeployCmd()
	err := deploy.RunDeployCmd()

	if err != nil {
		log.Fatal(err)
	}

	err = initialize.RunInitCmd()
	if err != nil {
		log.Fatal(err)
	}

	err = local.RunLocalCmd()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
