package main

import (
	"fmt"

	mainframepkg "github.com/fouralarmfire/grootsay/mainframe"
	replicapkg "github.com/fouralarmfire/grootsay/replica"
)

func main() {
	replica := replicapkg.NewReplica()
	groot := mainframepkg.NewMainframe(replica)

	fmt.Print("\n")
	groot.Say()
	fmt.Print("\n")
}
