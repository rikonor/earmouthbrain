package body

import (
	"github.com/rikonor/people/person/ear"
	"github.com/rikonor/people/person/mouth"
)

type BrainlessHTTPToConsoleBody struct {
	Ear   *ear.HTTPEar
	Mouth *mouth.ConsoleMouth
}

func NewBrainlessHTTPToConsoleBody(port string) *BrainlessHTTPToConsoleBody {
	// Get Mouth
	m := mouth.NewConsoleMouth()

	// Get Ear
	e := ear.NewHTTPEar(m.Say, port)

	return &BrainlessHTTPToConsoleBody{
		Ear:   e,
		Mouth: m,
	}
}

type BrainlessConsoleToConsoleBody struct {
	Ear   *ear.ConsoleEar
	Mouth *mouth.ConsoleMouth
}

func NewBrainlessConsoleToConsoleBody() *BrainlessConsoleToConsoleBody {
	// Get Mouth
	m := mouth.NewConsoleMouth()

	// Get Ear
	e := ear.NewConsoleEar(m.Say)

	return &BrainlessConsoleToConsoleBody{
		Ear:   e,
		Mouth: m,
	}
}

type BrainlessFileToFileBody struct {
	Ear   *ear.FileEar
	Mouth *mouth.FileMouth
}

func NewBrainlessFileToFileBody(inputFilePath, outputFilePath string) *BrainlessFileToFileBody {
	m := mouth.NewFileMouth(outputFilePath)
	e := ear.NewFileEar(m.Say, inputFilePath)

	return &BrainlessFileToFileBody{
		Ear:   e,
		Mouth: m,
	}
}

type BrainlessHTTPToFileBody struct {
	Ear   *ear.HTTPEar
	Mouth *mouth.FileMouth
}

func NewBrainlessHTTPToFileBody(port, outputFilePath string) *BrainlessHTTPToFileBody {
	m := mouth.NewFileMouth(outputFilePath)
	e := ear.NewHTTPEar(m.Say, port)

	return &BrainlessHTTPToFileBody{
		Ear:   e,
		Mouth: m,
	}
}