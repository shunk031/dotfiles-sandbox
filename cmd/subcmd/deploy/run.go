package deploy

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/shunk031/dotfiles/cmd/common"
)

func deployDotFiles(dotFileType string) error {
	dotPath, err := common.GetDotPath()
	if err != nil {
		return err
	}

	dotFilesPath := filepath.Join(dotPath, ".files", dotFileType)
	files, err := ioutil.ReadDir(dotFilesPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		filename := file.Name()
		if !strings.HasPrefix(filename, ".") {
			return fmt.Errorf("%s is not dotfile.", filename)
		}
		srcDotFilePath := filepath.Join(dotFilesPath, filename)
		dstDotFilePath := filepath.Join(os.Getenv("HOME"), filename)

		msg := fmt.Sprintf("%s -> %s", dstDotFilePath, srcDotFilePath)
		cmd := fmt.Sprintf("ln -sfnv %s %s", srcDotFilePath, dstDotFilePath)
		if err := common.Execute(msg, cmd); err != nil {
			return err
		}
	}
	return nil
}

func runDeploy(dotFileType string) error {
	if err := deployDotFiles(dotFileType); err != nil {
		return err
	}
	if err := common.CreateSymlinkHomeBinDir(); err != nil {
		return err
	}
	return nil
}

func RunDeployCmd() error {

	common.PrintInPurple("\n• Create symbolic links\n")

	if err := runDeployCommon(); err != nil {
		return err
	}
	return runDeploySystem()
}
