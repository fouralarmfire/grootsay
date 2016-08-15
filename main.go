package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/fouralarmfire/grootsay/ascii"
)

const maxLineLength = 80

func main() {
	fmt.Print("\n")
	if !(processStdin() || processArgs()) {
		ascii.SayIAmGroot()
	}
	ascii.CreateGroot()
}

func processStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdinMultiLineBubble()
		return true
	}
	return false
}

func processArgs() bool {
	args := os.Args[1:]
	if len(args) > 0 {
		argsSpeak(args)
		return true
	}
	return false
}

func argsSpeak(args []string) {
	text := strings.Join(args, " ")
	lineLength := len(text)
	if lineLength <= maxLineLength {
		oneLineBubble(text, lineLength)
	}
}

func oneLineBubble(text string, lineLength int) {
	fmt.Printf(color.MagentaString("         %s\n", topLine(lineLength)))
	fmt.Printf(color.CyanString("       %s  %s  %s\n", delimeters("only", 0), text, delimeters("only", 1)))
	fmt.Printf(color.MagentaString("         %s\n", bottomLine(lineLength)))
}

func stdinMultiLineBubble() {
	var del0, del1 string
	stdinScanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(color.MagentaString("   %s\n", topLine(maxLineLength)))
	for i := 0; stdinScanner.Scan(); i++ {
		if i == 0 {
			del0 = delimeters("first", 0)
			del1 = delimeters("first", 1)
		} else {
			del0 = delimeters("middle", 0)
			del1 = delimeters("middle", 1)
		}
		fmt.Printf(color.CyanString(" %s  %s %s\n", del0, pad(stdinScanner.Text()), del1))
	}
	fmt.Printf(color.MagentaString(" %s%s%s\n", delimeters("last", 0), bottomLine(maxLineLength), delimeters("last", 1)))
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
