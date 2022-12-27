package initialize

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestRunSshKeygen(t *testing.T) {
	if err := runSshKeygen(); err != nil {
		t.Fatal(err)
	}

	keyFilePath := filepath.Join(os.Getenv("HOME"), ".ssh", "id_ed25519")
	if _, err := os.Stat(keyFilePath); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("No key file for ssh is found to %s", keyFilePath)
	}
}
