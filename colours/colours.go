package colours

import (
	"github.com/fatih/color"
)

func PrintRed(str string) {
	red := color.New(color.FgRed).PrintlnFunc()
	red(str)
}

func PrintGreen(str string) {
	green := color.New(color.FgGreen).PrintlnFunc()
	green(str)
}

func PrintYellow(str string) {
	yellow := color.New(color.FgYellow).PrintlnFunc()
	yellow(str)
}

func PrintBlue(str string) {
	blue := color.New(color.FgBlue).PrintlnFunc()
	blue(str)
}

func PrintCyan(str string) {
	cyan := color.New(color.FgCyan).PrintlnFunc()
	cyan(str)
}

func PrintMagenta(str string) {
	magenta := color.New(color.FgMagenta).PrintlnFunc()
	magenta(str)
}
