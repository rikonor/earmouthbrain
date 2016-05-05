package earmouthbrain

import (
	"fmt"
)

type Ear struct {
	MessageHandlers []MessageHandler
}

func (e *Ear) RegisterMessageHandler(msgHandler MessageHandler) {
	e.MessageHandlers = append(e.MessageHandlers, msgHandler)
}

func (e *Ear) RelayMessage(msg Message) {
	if len(e.MessageHandlers) == 0 {
		fmt.Println("Ear can't relay message since no message handler was defined")
		return
	}

	for _, msgHandler := range e.MessageHandlers {
		msgHandler(msg)
	}
}
