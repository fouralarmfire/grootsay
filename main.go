package main

import (
	"fmt"

	"github.com/fouralarmfire/grootsay/ascii"
	"github.com/fouralarmfire/grootsay/creator"
	"github.com/fouralarmfire/grootsay/processor"
	"github.com/fouralarmfire/grootsay/screen"
)

func main() {
	screen := screen.NewScreen()
	image := ascii.NewImage()
	bubble := ascii.NewBubble()
	processor := processor.NewTextProcessor(screen.Width())
	grootsay := creator.NewGroot(processor, image, bubble)

	fmt.Print("\n")
	grootsay.Speak()
	fmt.Print("\n")
}
