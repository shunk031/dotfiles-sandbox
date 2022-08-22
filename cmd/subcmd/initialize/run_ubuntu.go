//+build ubuntu

package initialize

import "log"

func RunInitCmd() error {
	err := runInitMacOsCommon()
	if err != nil {
		log.Fatal(err)
	}
	return runInitMacOsArch()
}
