package processor

import (
	"bufio"
	"os"
	"strings"

	"github.com/fouralarmfire/grootsay/ascii"
	"github.com/fouralarmfire/grootsay/screen"
)

func ProcessStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		lines := getStdin()
		if len(lines) == 1 {
			ascii.OneLineBubble(lines[0], len(lines[0]))
		} else {
			ascii.StdinMultiLineBubble(lines)
		}
		return true
	}
	return false
}

func ProcessArgs() bool {
	args := os.Args[1:]
	sw := screen.GetWidth()
	text := strings.Join(args, " ")
	if len(args) > 0 && len(text) < int(sw)-13 {
		ascii.OneLineBubble(text, len(text))
		return true
	}
	return false
}

func getStdin() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
