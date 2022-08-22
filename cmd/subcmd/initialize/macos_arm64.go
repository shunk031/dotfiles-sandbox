//+build arm64

package initialize

import "fmt"

// https://qiita.com/ueokande/items/fac0d1219dbbc8f44db7 みたいなことをやって、アーキテクチャごとに変える

func runInitMacOsArch() error {
	fmt.Println("Run initialization for MacOS arm64")
	return nil
}
