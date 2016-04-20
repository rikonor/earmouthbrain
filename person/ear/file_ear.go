package ear

import (
	"fmt"
	"os"

	"github.com/hpcloud/tail"
	"github.com/rikonor/people/person/dto"
)

// FileEar - follows a file and treats lines as messages
type FileEar struct {
	MessageHandler dto.MessageHandler
	FilePath       string
}

// NewFileEar - Create a new FileEar
func NewFileEar(msgHandler dto.MessageHandler, filePath string) *FileEar {
	if _, err := os.Stat(filePath); err != nil {
		if _, err := os.Create(filePath); err != nil {
			panic(fmt.Sprintf("Failed to create file: %s", filePath))
		}
	}

	fe := FileEar{
		MessageHandler: msgHandler,
		FilePath:       filePath,
	}

	go fe.Listen()

	return &fe
}

func (fe *FileEar) relayMessage(msg dto.Message) {
	if fe.MessageHandler == nil {
		fmt.Println("Ear can't relay message sinfe no message handler was defined")
		return
	}
	fe.MessageHandler(msg)
}

func (fe *FileEar) Listen() {
	t, err := tail.TailFile(fe.FilePath, tail.Config{Follow: true})
	if err != nil {
		panic(fmt.Sprintf("Failed to tail file: %s", fe.FilePath))
	}

	for msgText := range t.Lines {
		msg := dto.StringToMessage(msgText.Text)

		fe.relayMessage(msg)
	}
}
