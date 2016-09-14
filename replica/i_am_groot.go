package replica

import (
	"github.com/fouralarmfire/grootsay/colours"
)

type Replica struct{}

func NewReplica() *Replica {
	return &Replica{}
}

func (i *Replica) Print() {
	colours.PrintRed("                 \\           ,=      Z8ZO")
	colours.PrintRed("                  \\          7=  D  77$7$Z:  OZZ")
	colours.PrintRed("                     O8      Z    $?$ D88Z, Z$Z,")
	colours.PrintCyan("                    8D$     $+     ZO  IID$?7$ND,")
	colours.PrintCyan("                     +OI=  ?$I?   78?~7IIMO?I8Z")
	colours.PrintCyan("                       ZO87II$??O$7I???I???IIO")
	colours.PrintBlue("                        O$77I$??I7II$I?I?O?II")
	colours.PrintBlue("                         Z$$7O7$??????I7$III7")
	colours.PrintMagenta("                         $7ZMMNM$$???$MMMNZ$7")
	colours.PrintMagenta("                         ZZMMMMMM7I??MMMMMM7$")
	colours.PrintMagenta("                         $77NMNM$7???7MMNM&Z$")
	colours.PrintBlue("                          77I??$7I$II?7$I?I7")
	colours.PrintBlue("                           ~$7$III87?II77$,")
	colours.PrintGreen("                                 `NDO'        8")
	colours.PrintGreen("                               OO$ZZZ78I    $ZO")
	colours.PrintGreen("                            8IZO8$7?7O?IIOZOO")
	colours.PrintGreen("                          $O$    .II?.  'O")
	colours.PrintGreen("                         $'       ZID")
	colours.PrintGreen("                                 ,7Z$,")
	colours.PrintYellow("                            +8NMO7ZO77DZI:")
	colours.PrintYellow("                           ++?$8OOO8N8DMN7+")
	colours.PrintYellow("                           ?+++=~~====+=++?")
	colours.PrintYellow("                           ??++=~~=++++++??")
	colours.PrintYellow("                           ???+=~~=++++++??")
}

func (i *Replica) DefaultMessage() {
	colours.PrintMagenta("        ____________")
	colours.PrintCyan("      <  I AM GROOT  >")
	colours.PrintMagenta("        ------------")
}
