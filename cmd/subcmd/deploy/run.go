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

func runDeployCommon() error {
	return runDeploy("common")
}

func runDeployClient() error {
	return runDeploy("client")
}

func runDeployServer() error {
	return runDeploy("server")
}

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
		err := Execute(msg, cmd)

		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func RunDeployCmd() error {
	if len(os.Args) < 2 {
		fmt.Println("Expected to specify client or server")
		os.Exit(1)
	}

	printInPurple("\n• Create symbolic links\n")

	err := runDeployCommon()
	if err != nil {
		log.Fatal(err)
	}

	machine := os.Args[2]
	if machine == "client" {
		return runDeployClient()
	} else if machine == "server" {
		return runDeployServer()
	} else {
		return fmt.Errorf("Invalid type: %s", machine)
	}
}
