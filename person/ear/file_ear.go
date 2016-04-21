package ear

import (
	"fmt"
	"os"

	"github.com/hpcloud/tail"
	"github.com/rikonor/people/person/dto"
)

// FileEar - follows a file and treats lines as messages
type FileEar struct {
	Ear
	FilePath string
}

// NewFileEar - Create a new FileEar
func NewFileEar(msgHandler dto.MessageHandler, filePath string) *FileEar {
	if _, err := os.Stat(filePath); err != nil {
		if _, err := os.Create(filePath); err != nil {
			panic(fmt.Sprintf("Failed to create file: %s", filePath))
		}
	}

	fe := FileEar{
		FilePath: filePath,
	}
	fe.MessageHandler = msgHandler

	go fe.Listen()

	return &fe
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
