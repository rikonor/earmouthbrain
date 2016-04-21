package main

import (
	"fmt"

	"github.com/rikonor/people/person/dto"
	"github.com/rikonor/people/person/ear"
)

func msgHandler(msg dto.Message) {
	fmt.Println("Got a special message:", msg)
}

func ExampleFileEar() {
	fe := ear.NewFileEar("./testfile1", "./testfile2")
	fe.RegisterMessageHandler(msgHandler)

	var x string
	fmt.Scanln(&x)
}

func ExampleConsoleEar() {
	ce := ear.NewConsoleEar()
	ce.RegisterMessageHandler(msgHandler)
	ce.Listen()

	// ce.Listen blocks (because it has to read from stdin)
}

func ExampleHTTPEar() {
	he := ear.NewHTTPEar("8080")
	he.RegisterMessageHandler(msgHandler)

	var x string
	fmt.Scanln(&x)
}

func main() {
	ExampleFileEar()
}
