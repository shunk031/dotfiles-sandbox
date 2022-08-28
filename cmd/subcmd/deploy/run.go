package deploy

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shunk031/dotfiles/cmd/common"
)

func runDeploy(dotFileType string) error {

	dotPath, err := common.GetDotPath()
	if err != nil {
		log.Fatal(err)
	}

	dotFilesPath := filepath.Join(dotPath, ".files", dotFileType)
	files, err := ioutil.ReadDir(dotFilesPath)
	if err != nil {
		log.Fatal(err)
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
		err := common.Execute(msg, cmd)

		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func RunDeployCmd() error {

	common.PrintInPurple("\n• Create symbolic links\n")

	err := runDeployCommon()
	if err != nil {
		log.Fatal(err)
	}
	return runDeploySystem()
}
