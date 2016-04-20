package ear

import (
	"fmt"

	"github.com/rikonor/people/person/dto"
)

// ConsoleEar - listens to input from the console
type ConsoleEar struct {
	MessageHandler dto.MessageHandler
	// once determines whether the ear will only listen to one message
	once bool
}

// NewConsoleEar - Create a new ConsoleEar
func NewConsoleEar(msgHandler dto.MessageHandler) *ConsoleEar {
	return &ConsoleEar{
		MessageHandler: msgHandler,
		once:           false,
	}
}

func (ce *ConsoleEar) relayMessage(msg dto.Message) {
	if ce.MessageHandler == nil {
		fmt.Println("Ear can't relay message since no message handler was defined")
		return
	}
	ce.MessageHandler(msg)
}

func (ce *ConsoleEar) Listen() {
	for {
		var msgText string
		fmt.Scanln(&msgText)

		msg := dto.StringToMessage(msgText)
		ce.relayMessage(msg)

		if ce.once {
			break
		}
	}
}
