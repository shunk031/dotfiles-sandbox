package initialize

import (
	"fmt"
	"log"
)

func runInitMacOs() error {
	err := runInitMacOsCommon()
	if err != nil {
		log.Fatal(err)
	}

	arch := getCpuArch()
	if arch == "x86_64" {
		return runInitMacOsAmd64()
	} else if arch == "arm64" {
		return runInitMacOsArm64()
	} else {
		return fmt.Errorf("Invalid CPU architecture: %s", arch)
	}
}

func RunInitCmd() error {
	os, err := getOs()
	if err != nil {
		log.Fatal(err)
	}

	if os == "macos" {
		return runInitMacOs()
	} else if os == "ubuntu" {
		return runInitUbuntu()
	} else {
		return fmt.Errorf("Invalid OS: %s", os)
	}
}
