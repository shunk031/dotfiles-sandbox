//go:build client

package deploy

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func TestRunDeploySystem(t *testing.T) {
	runDeployCommon()
	runDeploySystem()

	dotPath, err := common.GetDotPath()
	if err != nil {
		log.Fatal(err)
	}

	// get paths from .files/common
	commonDotFilesDirPath := filepath.Join(dotPath, ".files", "common")
	commonDotFiles, err := ioutil.ReadDir(commonDotFilesDirPath)
	if err != nil {
		log.Fatal(err)
	}
	var commonDotFilePaths []string
	for _, commonDotFilePath := range commonDotFiles {
		commonDotFilePaths = append(commonDotFilePaths, filepath.Join(commonDotFilesDirPath, commonDotFilePath.Name()))
	}
	numCommonDotFiles := len(commonDotFilePaths)
	actualNumCommonDotFiles := 3
	if numCommonDotFiles != actualNumCommonDotFiles {
		t.Fatalf("# of dotfiles for common is %d. But found %d.", actualNumCommonDotFiles, numCommonDotFiles)
	}

	// get paths from .files/client
	clientDotFilesDirPath := filepath.Join(dotPath, ".files", "client")
	clientDotFiles, err := ioutil.ReadDir(clientDotFilesDirPath)
	if err != nil {
		log.Fatal(err)
	}
	var clientDotFilePaths []string
	for _, clientDotFilePath := range clientDotFiles {
		clientDotFilePaths = append(clientDotFilePaths, filepath.Join(clientDotFilesDirPath, clientDotFilePath.Name()))
	}
	numClientDotFiles := len(clientDotFilePaths)
	actualNumClientDotFiles := 7
	if numClientDotFiles != actualNumClientDotFiles {
		t.Fatalf("# of dotfiles for client is %d. But found %d.", actualNumClientDotFiles, numClientDotFiles)
	}

	numMyDotFiles := numCommonDotFiles + numClientDotFiles
	numActualMyDotFiles := actualNumCommonDotFiles + actualNumClientDotFiles
	if numMyDotFiles != 10 {
		t.Fatalf("# of dotfiles (common + client) is %d. But found %d", numActualMyDotFiles, numMyDotFiles)
	}

	// compare dotfiles in home directory
	homeFiles, err := ioutil.ReadDir(os.Getenv("HOME"))
	if err != nil {
		log.Fatal(err)
	}
	var homeDotFilePaths []string
	for _, homeFile := range homeFiles {

		// not directory & has dot prefix
		if !homeFile.IsDir() && strings.HasPrefix(homeFile.Name(), ".") {
			p := filepath.Join(os.Getenv("HOME"), homeFile.Name())
			info, err := os.Lstat(p)
			if err != nil {
				log.Fatal(err)
			}

			// is symbolic link?
			if info.Mode()&os.ModeSymlink == os.ModeSymlink {
				homeDotFilePaths = append(homeDotFilePaths, p)
			}
		}
	}

	// compare dotfiles in home directory and my dotfiles
	numHomeDotFiles := len(homeDotFilePaths)
	if numMyDotFiles != numHomeDotFiles {
		t.Fatalf("# of my dotfiles is %d. But found %d", numMyDotFiles, numHomeDotFiles)
	}
}
