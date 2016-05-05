package earmouthbrain

type Mouth struct {
	inputChannel               chan Message
	isWaitingForSomethingToSay bool
}

func (m *Mouth) Init(sayFunc MessageHandler) {
	m.inputChannel = make(chan Message, DefaultInputBufferCapacity)
	m.waitForSomethingToSay(sayFunc)
}

func (m *Mouth) Say(msg Message) {
	if m.isWaitingForSomethingToSay {
		m.inputChannel <- msg
	}
}

func (m *Mouth) waitForSomethingToSay(saySync MessageHandler) {
	// Start handling messages
	go func() {
		for {
			select {
			case msg := <-m.inputChannel:
				saySync(msg)
			}
		}
	}()

	m.isWaitingForSomethingToSay = true
}
