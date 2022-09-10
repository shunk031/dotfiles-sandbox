package initialize

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
	"github.com/shunk031/dotfiles/cmd/subcmd/initialize/app"
)

func runSshKeygen() error {
	common.PrintInPurple("\n   Setup for ssh key\n")

	email := "shunsuke.kitada.0831@gmail.com"
	outputKeyfile := filepath.Join(os.Getenv("HOME"), ".ssh", "id_ed25519")

	if !common.PathExists(outputKeyfile) {
		cmd := fmt.Sprintf("ssh-keygen -q -t ed25519 -C \"%s\" -N \"\" -f %s <<<y >/dev/null 2>&1", email, outputKeyfile)
		msg := fmt.Sprintf("Generating a new SSH key to %s", outputKeyfile)
		return common.Execute(msg, cmd)
	} else {
		msg := fmt.Sprintf("File %s already exists", outputKeyfile)
		common.PrintSuccess(msg)
		return nil
	}
}

func installPowerlevel10k() error {
	dir := filepath.Join(os.Getenv("HOME"), ".zprezto", "modules", "prompt", "external", "powerlevel10k")
	url := "https://github.com/romkatv/powerlevel10k.git"

	if err := common.RemoveDir(dir); err != nil {
		return err
	}
	return common.Execute(
		fmt.Sprintf("Clone to %s", dir),
		fmt.Sprintf("git clone %s %s", url, dir),
	)
}

func InstallPowerlevel10k() error {
	common.PrintInPurple("\n   Install powerlevel10k\n")
	return installPowerlevel10k()
}

func runInitCommon() error {

	if err := runSshKeygen(); err != nil {
		return err
	}
	if err := InstallPowerlevel10k(); err != nil {
		return err
	}
	if err := app.InstallFzf(); err != nil {
		return err
	}
	return nil
}
