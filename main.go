package main

import (
	"fmt"

	collectorpkg "github.com/fouralarmfire/grootsay/collector"
	mainframepkg "github.com/fouralarmfire/grootsay/mainframe"
	replicapkg "github.com/fouralarmfire/grootsay/replica"
	screenpkg "github.com/fouralarmfire/grootsay/screen"
)

func main() {
	screen := screenpkg.NewScreen()
	replica := replicapkg.NewReplica()
	bubble := replicapkg.NewBubble()
	collector := collectorpkg.NewTextCollector(screen.Width())
	groot := mainframepkg.NewMainframe(collector, replica, bubble)

	fmt.Print("\n")
	groot.Speak()
	fmt.Print("\n")
}
