package main

import (
	"fmt"

	"github.com/rikonor/people/person/body"
)

func Example1() {
	body.NewBrainlessHTTPToConsoleBody("8080")

	var x string
	fmt.Scanln(&x)
}

func Example2() {
	b := body.NewBrainlessConsoleToConsoleBody()
	b.Ear.Listen()
}

func Example3() {
	body.NewBrainlessFileToFileBody("./input", "./output")

	var x string
	fmt.Scanln(&x)
}

func Example4() {
	body.NewBrainlessHTTPToFileBody("8080", "./output")

	var x string
	fmt.Scanln(&x)
}

func main() {
	// Example1()
	// Example2()
	// Example3()
	Example4()
}
