package mouth

import "github.com/rikonor/people/person/dto"

type Mouth interface {
	Say(dto.Message)
}
