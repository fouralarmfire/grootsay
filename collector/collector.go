package collector

import (
	"bufio"
	"os"
	"strings"
)

type TextCollector struct {
	screenWidth int
	lines       []string
}

func NewTextCollector(sw int) *TextCollector {
	return &TextCollector{
		screenWidth: sw,
	}
}

func (c *TextCollector) ReceivedInput() (bool, []string) {
	if (c.receivedStdin() || c.receivedArgs()) && c.checkLines() {
		return true, c.lines
	}
	return false, []string{}
}

func (c *TextCollector) receivedStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			c.lines = append(c.lines, scanner.Text())
		}
		return true
	}

	return false
}

func (c *TextCollector) receivedArgs() bool {
	args := os.Args[1:]
	text := strings.Join(args, " ")
	if len(args) > 0 && len(text) < int(c.screenWidth)-13 {
		c.lines = append(c.lines, text)
		return true
	}
	return false
}

func (c *TextCollector) checkLines() bool {
	if len(c.lines) == 0 {
		return false
	}
	return true
}
