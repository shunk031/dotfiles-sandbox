//go:build darwin

package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/shunk031/dotfiles/cmd/common"
)

func TestInstallLima(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}

	if err := installLima(); err != nil {
		t.Fatal(err)
	}
	if !common.CmdExists("lima") || !common.CmdExists("limatcl") {
		t.Fatal("Command not found: lima")
	}
}

func TestInstallDockerCli(t *testing.T) {
	if err := installDockerCli(); err != nil {
		t.Fatal(err)
	}
	if !common.CmdExists("docker") {
		t.Fatal("Command not found: docker")
	}
}

func TestInstallDockerCompose(t *testing.T) {
	if err := installDockerCompose(); err != nil {
		t.Fatal(err)
	}
	if !common.CmdExists("docker-compose") {
		t.Fatal("Command not found: docker-compose")
	}
}

func TestInstallDocker(t *testing.T) {
	if common.CmdExists("lima") || common.CmdExists("limactl") {
		if err := common.BrewUninstall("lima", "lima"); err != nil {
			t.Fatal(err)
		}
	}
	if common.CmdExists("docker") {
		dockerCmdPath := filepath.Join(os.Getenv("HOME"), "bin", "docker")
		if err := os.Remove(dockerCmdPath); err != nil {
			t.Fatal(err)
		}
	}
	if common.CmdExists("docker-compose") {
		dockerComposeCmdPath := filepath.Join(os.Getenv("HOME"), "bin", "docker-compose")
		if err := os.Remove(dockerComposeCmdPath); err != nil {
			t.Fatal(err)
		}
	}

	if err := InstallDocker(); err != nil {
		t.Fatal(err)
	}
	if !common.CmdExists("lima") || !common.CmdExists("limatcl") {
		t.Fatal("Command not found: lima")
	}
	if !common.CmdExists("docker") {
		t.Fatal("Command not found: docker")
	}
	if !common.CmdExists("docker-compose") {
		t.Fatal("Command not found: docker-compose")
	}
}
