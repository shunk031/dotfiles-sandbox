package common

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"gopkg.in/ini.v1"
)

const (
	SHELL      = "/bin/bash"
	maxCharLen = 170
)

func isDebug() bool {
	flag := os.Getenv("DEBUG_DOTFILES")
	if flag != "" {
		return true
	} else {
		return false
	}
}

func GetDotPath() (string, error) {
	dotPath := os.Getenv("DOTPATH")
	if dotPath != "" {
		return dotPath, nil
	} else {
		return "", fmt.Errorf("Need to specify DOTPATH")
	}
}

func Mkd(p string) {
	if len(p) > 0 {
		fInfo, err := os.Stat(p)
		if err != nil {
			log.Fatal(err)
		} else if errors.Is(err, os.ErrNotExist) {
			msg := p
			cmd := fmt.Sprintf("mkdir -p %s", p)
			err := Execute(msg, cmd)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if !!fInfo.IsDir() {
				printError(fmt.Errorf("%s - a file with the same name already exists", p))
			}
		}
	}
}

func ExecuteCmd(cmd string) error {

	c := exec.Command(SHELL, "-c", cmd)

	var out bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &out
	c.Stderr = &stderr

	err := c.Run()
	if err != nil {
		return errors.New(fmt.Sprintf("command %s; %s", cmd, fmt.Sprint(err)))
	} else {
		return nil
	}
}

func Execute(msg string, cmd string) error {
	// ref. https://gist.github.com/bamoo456/7e21773e8ef742a726c041f5f0019d2e

	// settings for the spiner
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = "  ["

	if len(cmd) > maxCharLen {
		s.Suffix = fmt.Sprintf("] %s", cmd[:maxCharLen])
	} else {
		s.Suffix = fmt.Sprintf("] %s", cmd)
	}

	// after the settings, start the spiner
	s.Start()

	// build the command
	c := exec.Command(SHELL, "-c", cmd)
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}

	// run the command
	if err := c.Start(); err != nil {
		return err
	}

	// for handling stdout
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			txt := scanner.Text()
			if len(txt) > maxCharLen {
				txt = txt[:maxCharLen]
			}
			s.Suffix = fmt.Sprintf("] %s", txt)
		}
	}()

	// for handling stderr
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			txt := scanner.Text()
			if len(txt) > maxCharLen {
				txt = txt[:maxCharLen]
			}
			s.Suffix = fmt.Sprintf("] %s", txt)
		}
	}()

	// waiting for finishing the command
	if err := c.Wait(); err != nil {
		return err
	}

	// stop the spiner
	s.Stop()

	// print the result
	printResult(msg, err)

	return err
}

func PathExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func CmdExists(c string) bool {
	cmd := exec.Command("/bin/bash", "-c", "command", "-v", c)
	if err := cmd.Run(); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Extract(archive string, outputDir string) error {

	if CmdExists("tar") {
		msg := fmt.Sprintf("Extract from %s to %s", archive, outputDir)
		cmd := fmt.Sprintf("tar -zxf %s --strip-components 1 -C %s", archive, outputDir)
		return Execute(msg, cmd)
	} else {
		return fmt.Errorf("Command not found: tar")
	}

}

func readOsRelease(cfgFile string) map[string]string {
	cfg, err := ini.Load(cfgFile)
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	cfgParams := make(map[string]string)
	cfgParams["ID"] = cfg.Section("").Key("ID").String()
	return cfgParams

}

func getLinuxDistribution() string {
	osInfo := readOsRelease("/etc/os-release")
	return osInfo["ID"]
}

func GetOs() (string, error) {
	kernelName := runtime.GOOS
	if kernelName == "darwin" {
		return "macos", nil
	} else if kernelName == "linux" {
		return getLinuxDistribution(), nil
	} else {
		return kernelName, fmt.Errorf("Invalid OS: %s", kernelName)
	}
}

func GetCpuArch() string {
	out, err := exec.Command("uname", "-m").Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(string(out), "\n")
}

func RemoveDir(dir string) error {
	cmd := fmt.Sprintf("rm -rf %s", dir)
	msg := fmt.Sprintf("Remove %s", dir)
	return Execute(msg, cmd)
}
