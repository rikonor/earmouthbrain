package brain

import "github.com/rikonor/people/person/dto"

type Brain interface {
	ProcessMessage(dto.Message)
}

type RepeatingBrain struct {
}

func (rb *RepeatingBrain) ProcessMessage(dto.Message) {

}
