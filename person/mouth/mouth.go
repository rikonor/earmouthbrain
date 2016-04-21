package mouth

import "github.com/rikonor/people/person/dto"

const DefaultInputBufferCapacity = 50

type Mouth struct {
	InputChannel               chan dto.Message
	isWaitingForSomethingToSay bool
}

func (m *Mouth) SayAsync(msg dto.Message) {
	if m.isWaitingForSomethingToSay {
		m.InputChannel <- msg
	}
}

func (m *Mouth) waitForSomethingToSay(saySync dto.MessageHandler) {
	m.isWaitingForSomethingToSay = true

	for {
		select {
		case msg := <-m.InputChannel:
			saySync(msg)
		}
	}
}
