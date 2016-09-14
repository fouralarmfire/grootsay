package mainframe

type Collector interface {
	ReceivedInput() (bool, []string)
}

type Replica interface {
	DefaultMessage()
	Print()
}

type Bubble interface {
	CustomMessage(text []string)
}

type Mainframe struct {
	bubble    Bubble
	collector Collector
	replica   Replica
}

func NewMainframe(c Collector, r Replica, b Bubble) *Mainframe {
	return &Mainframe{
		bubble:    b,
		collector: c,
		replica:   r,
	}
}

func (m *Mainframe) Speak() {
	rec, text := m.collector.ReceivedInput()
	if rec {
		m.bubble.CustomMessage(text)
	} else {
		m.replica.DefaultMessage()
	}
	m.replica.Print()
}
