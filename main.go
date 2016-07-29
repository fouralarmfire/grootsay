package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	fmt.Print("\n")
	if (processStdin() || processArgs()) != true {
		bubble, _ := ioutil.ReadFile(path.Join("ascii/i_am_groot"))
		fmt.Print(string(bubble))
	}
	createGroot()
}

func processStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdinScanner := bufio.NewScanner(os.Stdin)
		for stdinScanner.Scan() {
			fmt.Println(stdinScanner.Text())
		}
		return true
	} else {
		return false
	}
}

func processArgs() bool {
	args := os.Args[1:]
	if len(args) > 0 {
		fmt.Printf("         %s\n", strings.Trim(fmt.Sprint(args), "[]"))
		return true
	}
	return false
}

func createGroot() {
	groot, _ := ioutil.ReadFile("ascii/minigroot")
	fmt.Print(string(groot))
}
