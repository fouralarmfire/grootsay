package processor

import (
	"bufio"
	"os"
	"strings"
)

type TextProcessor struct {
	screenWidth int
	lines       []string
}

func NewTextProcessor(sw int) *TextProcessor {
	return &TextProcessor{
		screenWidth: sw,
	}
}

func (p *TextProcessor) ReceivedInput() (bool, []string) {
	if (p.receivedStdin() || p.receivedArgs()) && p.checkLines() {
		return true, p.lines
	}
	return false, []string{}
}

func (p *TextProcessor) receivedStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			p.lines = append(p.lines, scanner.Text())
		}
		return true
	}

	return false
}

func (p *TextProcessor) receivedArgs() bool {
	args := os.Args[1:]
	text := strings.Join(args, " ")
	if len(args) > 0 && len(text) < int(p.screenWidth)-13 {
		p.lines = append(p.lines, text)
		return true
	}
	return false
}

func (p *TextProcessor) checkLines() bool {
	if len(p.lines) == 0 {
		return false
	}
	return true
}
