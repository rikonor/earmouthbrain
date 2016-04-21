package ear

import (
	"fmt"

	"github.com/rikonor/people/person/dto"
)

type Ear struct {
	MessageHandlers []dto.MessageHandler
}

func (e *Ear) RegisterMessageHandler(msgHandler dto.MessageHandler) {
	e.MessageHandlers = append(e.MessageHandlers, msgHandler)
}

func (e *Ear) relayMessage(msg dto.Message) {
	if len(e.MessageHandlers) == 0 {
		fmt.Println("Ear can't relay message since no message handler was defined")
		return
	}

	for _, msgHandler := range e.MessageHandlers {
		msgHandler(msg)
	}
}
