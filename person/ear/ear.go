package ear

import "github.com/rikonor/people/person/dto"

type Ear interface {
	Listen()
	relayMessage(dto.Message)
}
