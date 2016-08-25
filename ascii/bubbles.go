package ascii

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const maxLineLength = 80

func StdinMultiLineBubble(lines []string) {
	var del0, del1 string
	fmt.Printf(color.MagentaString("   %s\n", topLine(maxLineLength)))
	for i, line := range lines {
		if i == 0 {
			del0 = delimeters("first", 0)
			del1 = delimeters("first", 1)
		} else if i == len(lines)-1 {
			del0 = delimeters("last", 0)
			del1 = delimeters("last", 1)
		} else {
			del0 = delimeters("middle", 0)
			del1 = delimeters("middle", 1)
		}
		fmt.Printf(color.CyanString(" %s  %s %s\n", del0, pad(strings.TrimSpace(line)), del1))
	}
	fmt.Printf(color.MagentaString("   %s\n", bottomLine(maxLineLength)))
}

func ArgsSpeak(args []string) {
	text := strings.Join(args, " ")
	lineLength := len(text)
	if lineLength <= maxLineLength {
		OneLineBubble(text, lineLength)
	}
}

func OneLineBubble(text string, lineLength int) {
	fmt.Printf(color.MagentaString("         %s\n", topLine(lineLength)))
	fmt.Printf(color.CyanString("       %s  %s  %s\n", delimeters("only", 0), text, delimeters("only", 1)))
	fmt.Printf(color.MagentaString("         %s\n", bottomLine(lineLength)))
}

func delimeters(loc string, index int) string {
	var dels = map[string][]string{
		"first":  []string{"/", "\\"},
		"middle": []string{"|", "|"},
		"last":   []string{"\\", "/"},
		"only":   []string{"<", ">"},
	}
	return dels[loc][index]
}

func pad(text string) string {
	padding := strings.Repeat(" ", maxLineLength-len(text))
	return text + padding
}

func topLine(length int) string {
	return strings.Repeat("_", length+2)
}

func bottomLine(length int) string {
	return strings.Repeat("-", length+2)
}
