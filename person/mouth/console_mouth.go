package mouth

import (
	"fmt"

	"github.com/rikonor/people/person/dto"
)

type ConsoleMouth struct {
}

func NewConsoleMouth() *ConsoleMouth {
	return &ConsoleMouth{}
}

func (m *ConsoleMouth) Say(msg dto.Message) {
	fmt.Println("I have this to say:", msg)
}
