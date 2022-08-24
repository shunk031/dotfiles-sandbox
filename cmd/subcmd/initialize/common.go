package initialize

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
	"github.com/shunk031/dotfiles/cmd/subcmd/initialize/app"
)

func runSshKeygen() {
	email := "shunsuke.kitada.0831@gmail.com"
	outputKeyfile := filepath.Join(os.Getenv("HOME"), ".ssh", "id_ed25519")
	cmd := fmt.Sprintf("ssh-keygen -q -t ed25519 -C \"%s\" -N \"\" -f %s", email, outputKeyfile)
	msg := fmt.Sprintf("Generating a new SSH key to %s", outputKeyfile)
	common.Execute(msg, cmd)
}

func runInitCommon() {

	runSshKeygen()

	app.InstallFzf()
}
