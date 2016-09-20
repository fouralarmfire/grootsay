package colours

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func Randomize(str string, ind int) {
	color.Set(randomColour(ind))
	defer color.Unset()
	fmt.Println(str)
}

func randomColour(ind int) color.Attribute {
	var colours = []color.Attribute{
		color.FgRed,
		color.FgGreen,
		color.FgYellow,
		color.FgBlue,
		color.FgMagenta,
		color.FgCyan,
	}

	rand.Seed(int64(ind-(-1)) * time.Now().Unix())
	return colours[rand.Int()%len(colours)]
}

func PrintCyan(str string) {
	cyan := color.New(color.FgCyan).PrintlnFunc()
	cyan(str)
}

func PrintMagenta(str string) {
	magenta := color.New(color.FgMagenta).PrintlnFunc()
	magenta(str)
}
