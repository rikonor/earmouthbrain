package mouth

import (
	"fmt"
	"os"

	"github.com/rikonor/people/person/dto"
)

type FileMouth struct {
	Mouth
	FilePath string
	file     *os.File
}

func NewFileMouth(filePath string) *FileMouth {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open file for writing: %s", err))
	}

	// Notice file is never closed

	m := FileMouth{
		FilePath: filePath,
		file:     f,
	}
	m.InputChannel = make(chan dto.Message, DefaultInputBufferCapacity)

	go m.waitForSomethingToSay(m.SaySync)

	return &m
}

func (m *FileMouth) SaySync(msg dto.Message) {
	msgText := string(msg) + "\n"
	m.file.WriteString(msgText)
}
