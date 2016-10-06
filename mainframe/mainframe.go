package mainframe

import (
	"github.com/fouralarmfire/grootsay/bubbles"
	"github.com/fouralarmfire/grootsay/collector"
)

type Replica interface {
	DefaultMessage()
	Print()
}

type Mainframe struct {
	bubble    *bubbles.Bubble
	collector *collector.TextCollector
	replica   Replica
}

func NewMainframe(r Replica) *Mainframe {
	return &Mainframe{
		bubble:    bubbles.NewBubble(),
		collector: collector.NewTextCollector(),
		replica:   r,
	}
}

func (m *Mainframe) Say() {
	ok, text := m.collector.ReceivedInput()
	if ok {
		m.bubble.CustomMessage(text)
	} else {
		m.replica.DefaultMessage()
	}
	m.replica.Print()
}
