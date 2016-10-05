package main

import (
	"fmt"

	collectorpkg "github.com/fouralarmfire/grootsay/collector"
	mainframepkg "github.com/fouralarmfire/grootsay/mainframe"
	replicapkg "github.com/fouralarmfire/grootsay/replica"
)

func main() {
	replica := replicapkg.NewReplica()
	bubble := replicapkg.NewBubble()
	collector := collectorpkg.NewTextCollector()
	groot := mainframepkg.NewMainframe(collector, replica, bubble)

	fmt.Print("\n")
	groot.Say()
	fmt.Print("\n")
}
