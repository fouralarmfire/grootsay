package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/fouralarmfire/grootsay/ascii"
)

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
		stdinScanner := bufio.NewScanner(os.Stdin)
		for stdinScanner.Scan() {
			fmt.Println(color.MagentaString(stdinScanner.Text()))
		}
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
	maxLineLength := 38
	text := strings.Join(args, " ")
	lineLength := len(text)
	if lineLength <= maxLineLength {
		fmt.Printf(color.MagentaString("         %s\n", topLine(lineLength)))
		fmt.Printf(color.CyanString("       %s  %s  %s\n", delimeters("only", 0), text, delimeters("only", 1)))
		fmt.Printf(color.MagentaString("         %s\n", bottomLine(lineLength)))
	}
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

func pad(text string, length int) string {
	padding := strings.Repeat(" ", length-(length+1))
	return text + padding
}

func topLine(length int) string {
	return strings.Repeat("_", length+2)
}

func bottomLine(length int) string {
	return strings.Repeat("-", length+2)
}
