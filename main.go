package main

import (
	"fmt"

	"github.com/fouralarmfire/xsay"
)

func main() {
	groot := xsay.New("groot.txt", "I AM GROOT!")

	fmt.Print("\n")
	groot.Say()
	fmt.Print("\n")
}
