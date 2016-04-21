package ear

import (
	"fmt"

	"github.com/rikonor/people/person/dto"
)

// ConsoleEar - listens to input from the console
type ConsoleEar struct {
	Ear
}

// NewConsoleEar - Create a new ConsoleEar
func NewConsoleEar(msgHandler dto.MessageHandler) *ConsoleEar {
	ce := ConsoleEar{}
	ce.MessageHandler = msgHandler
	return &ce
}

func (ce *ConsoleEar) Listen() {
	for {
		var msgText string
		fmt.Scanln(&msgText)

		msg := dto.StringToMessage(msgText)
		ce.relayMessage(msg)
	}
}
