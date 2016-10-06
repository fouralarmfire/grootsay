package mainframe

import (
	"github.com/fouralarmfire/grootsay/bubbles"
	"github.com/fouralarmfire/grootsay/collector"
	"github.com/fouralarmfire/grootsay/replica"
)

type Replica interface {
	DefaultMessage()
	Print()
}

type Mainframe struct {
	bubble    *bubbles.Bubble
	collector *collector.TextCollector
	replica   *replica.Replica
}

func NewMainframe(asciiArt []string, defaultMessage string) *Mainframe {
	return &Mainframe{
		bubble:    bubbles.NewBubble(defaultMessage),
		collector: collector.NewTextCollector(),
		replica:   replica.NewReplica(asciiArt),
	}
}

func (m *Mainframe) Say() {
	ok, text := m.collector.ReceivedInput()
	if ok {
		m.bubble.CustomMessage(text)
	} else {
		m.bubble.DefaultMessage()
	}
	m.replica.Print()
}
