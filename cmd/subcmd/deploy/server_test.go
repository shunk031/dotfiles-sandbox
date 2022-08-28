//go:build server

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

	// get paths from .files/server
	serverDotFilesDirPath := filepath.Join(dotPath, ".files", "server")
	serverDotFiles, err := ioutil.ReadDir(serverDotFilesDirPath)
	if err != nil {
		log.Fatal(err)
	}
	var serverDotFilePaths []string
	for _, serverDotFilePath := range serverDotFiles {
		serverDotFilePaths = append(serverDotFilePaths, filepath.Join(serverDotFilesDirPath, serverDotFilePath.Name()))
	}
	numServerDotFiles := len(serverDotFilePaths)
	actualNumServerDotFiles := 1
	if numServerDotFiles != actualNumServerDotFiles {
		t.Fatalf("# of dotfiles for server is %d. But found %d.", actualNumServerDotFiles, numServerDotFiles)
	}

	numMyDotFiles := numCommonDotFiles + numServerDotFiles
	actualNumMyDotFiles := 4
	if numMyDotFiles != actualNumMyDotFiles {
		t.Fatalf("# of my dotfiles is %d. But found %d.", actualNumMyDotFiles, numMyDotFiles)
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
