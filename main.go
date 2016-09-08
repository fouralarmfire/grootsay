package main

import (
	"fmt"

	"github.com/fouralarmfire/grootsay/ascii"
	"github.com/fouralarmfire/grootsay/processor"
)

func main() {
	fmt.Print("\n")
	if !(processor.ProcessStdin() || processor.ProcessArgs()) {
		ascii.SayIAmGroot()
	}
	ascii.CreateGroot()
}
