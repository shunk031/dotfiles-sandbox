//go:build ubuntu

package app

import "github.com/shunk031/dotfiles/cmd/common"

func InstallGhq() {
	common.PrintInPurple("\n   Install ghq\n\n")

	msg := "ghq"
	cmd := "go install github.com/x-motemen/ghq@latest"
	common.Execute(msg, cmd)

	ghqDir := filepath.Join(os.Getenv("HOME"), "ghq")
	cmd := fmt.Sprintf("mkdir -p %s", ghqDir)
	common.ExecuteCmd(cmd)
}
