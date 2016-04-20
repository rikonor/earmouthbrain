package main

import (
	"fmt"

	"github.com/rikonor/people/person/dto"
	"github.com/rikonor/people/person/ear"
)

func msgHandler(msg dto.Message) {
	fmt.Println("Got a special message:", msg)
}

func main() {
	ear.NewFileEar(msgHandler, "./testfile")

	var x string
	fmt.Scanln(&x)
}
