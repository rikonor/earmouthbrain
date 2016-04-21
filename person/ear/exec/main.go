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
	ear.NewFileEar(msgHandler, "./testfile")

	var x string
	fmt.Scanln(&x)
}

func ExampleConsoleEar() {
	ce := ear.NewConsoleEar(msgHandler)
	ce.Listen()
}

func ExampleHTTPEar() {
	ear.NewHTTPEar(msgHandler, "8080")

	var x string
	fmt.Scanln(&x)
}

func main() {
	ExampleHTTPEar()
}
