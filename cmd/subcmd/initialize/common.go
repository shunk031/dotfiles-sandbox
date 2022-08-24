package initialize

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
	"github.com/shunk031/dotfiles/cmd/subcmd/initialize/app"
)

func runSshKeygen() {
	common.PrintInPurple("\n   Setup for ssh key\n")

	email := "shunsuke.kitada.0831@gmail.com"
	outputKeyfile := filepath.Join(os.Getenv("HOME"), ".ssh", "id_ed25519")
	cmd := fmt.Sprintf("ssh-keygen -q -t ed25519 -C \"%s\" -N \"\" -f %s", email, outputKeyfile)
	msg := fmt.Sprintf("Generating a new SSH key to %s", outputKeyfile)
	common.Execute(msg, cmd)
}

func installPowerlevel10k() error {
	dir := filepath.Join(os.Getenv("HOME"), ".zprezto", "modules", "prompt", "external", "powerlevel10k")
	url := "https://github.com/romkatv/powerlevel10k.git"
	cmd := fmt.Sprintf("rm -rf %s && git clone --quiet %s %s", dir, url, dir)
	msg := fmt.Sprintf("Clone to %s", dir)
	return common.Execute(msg, cmd)
}

func InstallPowerlevel10k() {
	common.PrintInPurple("\n   Install powerlevel10k\n")
	installPowerlevel10k()
}

func runInitCommon() {

	runSshKeygen()

	InstallPowerlevel10k()

	app.InstallFzf()
}
