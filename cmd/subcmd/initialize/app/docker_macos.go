//go:build darwin

package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/shunk031/dotfiles/cmd/common"
)

func installLima() error {
	// TODO (shunk031): run `limactl start docker.yaml`
	return common.BrewInstall("lima", "lima", common.BrewOpts{})
}

func downloadWithCurl(url string, f string) error {
	cmd := fmt.Sprintf("curl -LsSo %s %s", f, url)
	msg := fmt.Sprintf("Download %s to %s", url, f)
	return common.Execute(msg, cmd)
}

func moveToHomeBin(fp string, fname string) error {
	srcPath := filepath.Join(fp, fname)
	dstPath := filepath.Join(os.Getenv("HOME"), "bin", fname)
	return common.Execute(
		fmt.Sprintf("Move %s to %s", srcPath, dstPath),
		fmt.Sprintf("mv %s %s", srcPath, dstPath),
	)
}

func extractAndMoveTarFile(archive string) error {
	cmd := fmt.Sprintf("tar fvx %s", archive)
	msg := fmt.Sprintf("Extract from %s", archive)
	if err := common.Execute(msg, cmd); err != nil {
		return err
	}

	tgzBinPath := filepath.Join(archive, "docker")
	homeBinPath := filepath.Join(os.Getenv("HOME"), "bin", "docker")
	return moveToHomeBin(tgzBinPath, homeBinPath)
}

func installDockerCli() error {
	dir, err := ioutil.TempDir("", "docker-cli")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	cpu := common.GetCpuArch()
	ver := "20.10.10"
	url := fmt.Sprintf("https://download.docker.com/mac/static/stable/%s/docker-%s.tgz", cpu, ver)
	tgzFile := filepath.Join(dir, "docker.tgz")
	if err := downloadWithCurl(url, tgzFile); err != nil {
		return err
	}
	return extractAndMoveTarFile(tgzFile)
}

func installDockerCompose() error {
	dir, err := ioutil.TempDir("", "docker-compose")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	cpu := common.GetCpuArch()
	url := fmt.Sprintf("https://github.com/docker/compose/releases/latest/download/docker-compose-darwin-%s", cpu)
	execFile := filepath.Join(dir, "docker-compose")
	if err := downloadWithCurl(url, execFile); err != nil {
		return err
	}
	return moveToHomeBin(execFile, "docker-compose")
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
