package mouth

import (
	"fmt"

	"github.com/rikonor/people/person/dto"
)

type HTTPMouth struct {
}

func NewHTTPMouth() *HTTPMouth {
	return &HTTPMouth{}
}

func (m *HTTPMouth) Say(msg dto.Message) {
	fmt.Println("I have this to say:", msg)
}
