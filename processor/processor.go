package processor

import (
	"os"

	"github.com/fouralarmfire/grootsay/ascii"
)

func ProcessStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		ascii.StdinMultiLineBubble()
		return true
	}
	return false
}

func ProcessArgs() bool {
	args := os.Args[1:]
	if len(args) > 0 {
		ascii.ArgsSpeak(args)
		return true
	}
	return false
}
