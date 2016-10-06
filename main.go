package main

import (
	"fmt"

	mainframepkg "github.com/fouralarmfire/xsay/mainframe"
)

func main() {
	groot := mainframepkg.NewMainframe(grootAscii(), defaultMessage())

	fmt.Print("\n")
	groot.Say()
	fmt.Print("\n")
}
