package main

import (
	"strconv"
	"time"

	"github.com/rikonor/people/person/dto"
	"github.com/rikonor/people/person/mouth"
)

func ExampleConsoleMouth() {
	m := mouth.NewConsoleMouth()
	for i := 1; i < 10; i++ {
		m.InputChannel <- dto.StringToMessage(strconv.Itoa(i))
	}

	time.Sleep(1 * time.Second)
}

func ExampleFileMouth() {
	m := mouth.NewFileMouth("./testfile")
	for i := 1; i < 10; i++ {
		m.InputChannel <- dto.StringToMessage(strconv.Itoa(i))
	}

	time.Sleep(1 * time.Second)
}

func ExampleHTTPMouth() {
	m := mouth.NewHTTPMouth("localhost:8080")
	for i := 1; i < 10; i++ {
		m.InputChannel <- dto.StringToMessage(strconv.Itoa(i))
	}

	time.Sleep(1 * time.Second)
}

func main() {
	// ExampleConsoleMouth()
	// ExampleFileMouth()
	ExampleHTTPMouth()
}
