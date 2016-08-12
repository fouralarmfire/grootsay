package ascii

import (
	"github.com/fatih/color"
)

func CreateGroot() {
	printMagenta("        ___       ___")
	printMagenta("            \\,,,,/")
	printRed("                 \\           ,=      Z8ZO")
	printRed("                  \\          7=  D  77$7$Z:  OZZ")
	printRed("                     O8      Z    $?$ D88Z, Z$Z,")
	printCyan("                    8D$     $+     ZO  IID$?7$ND,")
	printCyan("                     +OI=  ?$I?   78?~7IIMO?I8Z")
	printCyan("                       ZO87II$??O$7I???I???IIO")
	printBlue("                        O$77I$??I7II$I?I?O?II")
	printBlue("                         Z$$7O7$??????I7$III7")
	printMagenta("                         $7ZMMNM$$???$MMMNZ$7")
	printMagenta("                         ZZMMMMMM7I??MMMMMM7$")
	printMagenta("                         $77NMNM$7???7MMNM&Z$")
	printBlue("                          77I??$7I$II?7$I?I7")
	printBlue("                           ~$7$III87?II77$,")
	printGreen("                                 `NDO'        8")
	printGreen("                               OO$ZZZ78I    $ZO")
	printGreen("                            8IZO8$7?7O?IIOZOO")
	printGreen("                          $O$    .II?.  'O")
	printGreen("                         $'       ZID")
	printGreen("                                 ,7Z$,")
	printYellow("                            +8NMO7ZO77DZI:")
	printYellow("                           ++?$8OOO8N8DMN7+")
	printYellow("                           ?+++=~~====+=++?")
	printYellow("                           ??++=~~=++++++??")
	printYellow("                           ???+=~~=++++++??")
}

func SayIAmGroot() {
	printMagenta("        _____________")
	printYellow("/```                     ```\\")
	printCyan("\\___      I AM GROOT     ___/")
}

func printRed(str string) {
	red := color.New(color.FgRed).PrintlnFunc()
	red(str)
}

func printGreen(str string) {
	green := color.New(color.FgGreen).PrintlnFunc()
	green(str)
}

func printYellow(str string) {
	yellow := color.New(color.FgYellow).PrintlnFunc()
	yellow(str)
}

func printBlue(str string) {
	blue := color.New(color.FgBlue).PrintlnFunc()
	blue(str)
}

func printCyan(str string) {
	cyan := color.New(color.FgCyan).PrintlnFunc()
	cyan(str)
}

func printMagenta(str string) {
	magenta := color.New(color.FgMagenta).PrintlnFunc()
	magenta(str)
}
