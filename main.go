package main

import (
	"fmt"

	"github.com/rikonor/people/person"
)

func main() {
	fmt.Println("Starting..")

	p := person.New("Or Rikon", 27)
	p.Listen(8080)

	var input string
	fmt.Scanln(&input)
}
