package ear

import (
	"fmt"

	"github.com/rikonor/people/person/dto"
)

type Ear struct {
	MessageHandler dto.MessageHandler
}

func (e *Ear) relayMessage(msg dto.Message) {
	if e.MessageHandler == nil {
		fmt.Println("Ear can't relay message since no message handler was defined")
		return
	}
	e.MessageHandler(msg)
}
