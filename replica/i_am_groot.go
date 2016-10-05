package replica

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

type Replica struct {
}

func NewReplica() *Replica {
	return &Replica{}
}

func (i *Replica) DefaultMessage() {
	color.Magenta("        ____________")
	color.Cyan("      <  I AM GROOT  >")
	color.Magenta("        ------------")
}

func (i *Replica) Print() {
	lines := []string{
		"                 \\           ,=      Z8ZO",
		"                  \\          7=  D  77$7$Z:  OZZ",
		"                     O8      Z    $?$ D88Z, Z$Z,",
		"                    8D$     $+     ZO  IID$?7$ND,",
		"                     +OI=  ?$I?   78?~7IIMO?I8Z",
		"                       ZO87II$??O$7I???I???IIO",
		"                        O$77I$??I7II$I?I?O?II",
		"                         Z$$7O7$??????I7$III7",
		"                         $7ZMMNM$$???$MMMNZ$7",
		"                         ZZMMMMMM7I??MMMMMM7$",
		"                         $77NMNM$7???7MMNM&Z$",
		"                          77I??$7I$II?7$I?I7",
		"                           ~$7$III87?II77$,",
		"                                 `NDO'        8",
		"                               OO$ZZZ78I    $ZO",
		"                            8IZO8$7?7O?IIOZOO",
		"                          $O$    .II?.  'O",
		"                         $'       ZID",
		"                                 ,7Z$,",
		"                            +8NMO7ZO77DZI:",
		"                           ++?$8OOO8N8DMN7+",
		"                           ?+++=~~====+=++?",
		"                           ??++=~~=++++++??",
		"                           ???+=~~=++++++??",
	}

	for i, line := range lines {
		randomize(line, i)
	}
}

func randomize(str string, ind int) {
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
