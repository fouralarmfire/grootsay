package creator

type Processor interface {
	ReceivedInput() (bool, []string)
}

type Image interface {
	DefaultMessage()
	CreateAscii()
}

type Bubble interface {
	CustomMessage(text []string)
}

type Groot struct {
	processor Processor
	image     Image
	bubble    Bubble
}

func NewGroot(p Processor, i Image, b Bubble) *Groot {
	return &Groot{
		processor: p,
		image:     i,
		bubble:    b,
	}
}

func (g *Groot) Speak() {
	rec, text := g.processor.ReceivedInput()
	if rec {
		g.bubble.CustomMessage(text)
	} else {
		g.image.DefaultMessage()
	}
	g.image.CreateAscii()
}
