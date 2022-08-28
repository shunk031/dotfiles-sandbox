//go:build client

package deploy

import "testing"

func TestRunDeploy(t *testing.T) {
	RunDeploy()

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
	actualNumClientDotFiles := 1
	if numClientDotFiles != actualNumClientDotFiles {
		t.Fatalf("# of dotfiles for client is %d. But found %d.", actualNumClientDotFiles, numClientDotFiles)
	}

	numMyDotFiles := numCommonDotFiles + numClientDotFiles
	if numMyDotFiles != 7 {
		t.Fatal("test failed")
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
		t.Fatal("test failed")
	}
}
