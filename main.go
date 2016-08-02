package main

import (
	"bufio"
	"fmt"
	"github.com/fouralarmfire/grootsay/ascii"
	"os"
	"strings"
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
			fmt.Println(stdinScanner.Text())
		}
		return true
	}
	return false
}

func processArgs() bool {
	args := os.Args[1:]
	if len(args) > 0 {
		fmt.Printf("         %s\n", strings.Trim(fmt.Sprint(args), "[]"))
		return true
	}
	return false
}
