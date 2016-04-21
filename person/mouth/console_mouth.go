package mouth

import (
	"fmt"

	"github.com/rikonor/people/person/dto"
)

type ConsoleMouth struct {
	Mouth
}

func NewConsoleMouth() *ConsoleMouth {
	m := ConsoleMouth{}
	m.InputChannel = make(chan dto.Message, DefaultInputBufferCapacity)
	go m.waitForSomethingToSay(m.SaySync)
	return &m
}

func (m *ConsoleMouth) SaySync(msg dto.Message) {
	fmt.Println("I have this to say:", msg)
}
