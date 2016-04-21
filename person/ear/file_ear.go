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
	FilePaths []string
}

// NewFileEar - Create a new FileEar
func NewFileEar(filePaths ...string) *FileEar {
	for _, filePath := range filePaths {
		if _, err := os.Stat(filePath); err != nil {
			if _, err := os.Create(filePath); err != nil {
				panic(fmt.Sprintf("Failed to create file: %s", filePath))
			}
		}
	}

	fe := FileEar{
		FilePaths: filePaths,
	}

	go fe.Listen()

	return &fe
}

func (fe *FileEar) Listen() {
	for _, filePath := range fe.FilePaths {
		go func(filePath string) {
			t, err := tail.TailFile(filePath, tail.Config{Follow: true})
			if err != nil {
				panic(fmt.Sprintf("Failed to tail file: %s", filePath))
			}

			for msgText := range t.Lines {
				msg := dto.StringToMessage(msgText.Text)

				fe.relayMessage(msg)
			}
		}(filePath)
	}
}
