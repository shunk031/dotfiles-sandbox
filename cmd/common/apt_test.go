//go:build linux

package common

import "testing"

func TestAptPackageIsInstalled(t *testing.T) {
	if AptPackageIsInstalled("package-does-not-exists") {
		t.Fatal("Test failed")
	}

	if !AptPackageIsInstalled("apt") {
		t.Fatal("Test failed")
	}
}

func TestAptInstall(t *testing.T) {
	if err := AptPurge("sl", "sl", AptOpts{}); err != nil {
		t.Fatal(err)
	}
	if err := AptInstall("sl", "sl", AptOpts{}); err != nil {
		t.Fatal(err)
	}

	if !AptPackageIsInstalled("sl") {
		t.Fatal("Test failed")
	}
}

func TestAptPurge(t *testing.T) {
	// test if the installed package
	if err := AptInstall("sl", "sl", AptOpts{}); err != nil {
		t.Fatal(err)
	}
	if err := AptPurge("sl", "sl", AptOpts{}); err != nil {
		t.Fatal(err)
	}
	if AptPackageIsInstalled("sl") {
		t.Fatal("Test failed")
	}

	// test if the purged package
	if err := AptPurge("sl", "sl", AptOpts{}); err != nil {
		t.Fatal(err)
	}
}
