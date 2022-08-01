package initialize

import "github.com/shunk031/dotfiles/cmd/common"

func runInitUbuntu() error {

	common.AptInstall("tmux", "tmux", "")
	common.AptInstall("tmux (pasteboard)", "xsel", "")
	common.AptInstall("cmake", "cmake", "")
	return nil

}
