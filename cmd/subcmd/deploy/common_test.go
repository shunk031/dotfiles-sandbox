package deploy

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func TestRunDeployCommon(t *testing.T) {
	runDeployCommon()

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
}
