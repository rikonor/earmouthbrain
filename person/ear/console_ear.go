package ear

import (
	"fmt"

	"github.com/rikonor/people/person/dto"
)

// ConsoleEar - listens to input from the console
type ConsoleEar struct {
	Ear
	handlers []dto.MessageHandler
}

// NewConsoleEar - Create a new ConsoleEar
func NewConsoleEar() *ConsoleEar {
	ce := ConsoleEar{}
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
