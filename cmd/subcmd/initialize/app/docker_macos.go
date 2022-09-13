//go:build darwin

package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

const (
	dockerCliVersion                 = "20.10.10"
	dockerDownloadURLTemplate        = "https://download.docker.com/mac/static/stable/%s/docker-%s.tgz"
	dockerComposeDownloadURLTemplate = "https://github.com/docker/compose/releases/latest/download/docker-compose-darwin-%s"
)

func installLima() error {
	// TODO (shunk031): run `limactl start docker.yaml`
	return common.BrewInstall("lima", "lima", common.BrewOpts{})
}

func extractAndMoveTarFile(archive string, basedir string) error {
	if err := common.Extract(archive, basedir); err != nil {
		return err
	}

	homeBinDir := filepath.Join(os.Getenv("HOME"), "bin")

	tgzBinPath := filepath.Join(basedir, "docker")
	homeBinPath := filepath.Join(homeBinDir, "docker")

	if !common.PathExists(homeBinDir) {
		if err := common.CreateSymlinkHomeBinDir(); err != nil {
			return err
		}
	}
	return os.Rename(tgzBinPath, homeBinPath)
}

func installDockerCli() error {
	dir, err := ioutil.TempDir("", "docker-cli")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	cpu := common.GetCpuArch()
	if cpu == "arm64" {
		cpu = "aarch64"
	}
	url := fmt.Sprintf(dockerDownloadURLTemplate, cpu, dockerCliVersion)
	tgzFile := filepath.Join(dir, "docker.tgz")
	if err := common.Download(url, tgzFile); err != nil {
		return err
	}
	return extractAndMoveTarFile(tgzFile, dir)
}

func installDockerCompose() error {
	dir, err := ioutil.TempDir("", "docker-compose")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	cpu := common.GetCpuArch()
	if cpu == "arm64" {
		cpu = "aarch64"
	}
	url := fmt.Sprintf(dockerComposeDownloadURLTemplate, cpu)
	execFile := filepath.Join(dir, "docker-compose")
	if err := common.Download(url, execFile); err != nil {
		return err
	}
	homeBinFile := filepath.Join(os.Getenv("HOME"), "bin", "docker-compose")
	if err := os.Rename(execFile, homeBinFile); err != nil {
		return err
	}
	return os.Chmod(homeBinFile, 0755)
}

func InstallDocker() error {
	common.PrintInPurple("\n   Docker\n")

	if err := installLima(); err != nil {
		return err
	}
	if err := installDockerCli(); err != nil {
		return err
	}
	if err := installDockerCompose(); err != nil {
		return err
	}
	return nil
}
