package main

import (
	"fmt"

	"github.com/rikonor/people/person/ear"
	"github.com/rikonor/people/person/mouth"
)

func main() {
	cm := mouth.NewConsoleMouth()
	ear.NewHTTPEar(cm.Say)

	var x string
	fmt.Scanln(&x)
}
