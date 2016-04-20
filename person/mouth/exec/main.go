package main

import "github.com/rikonor/people/person/mouth"

func ExampleConsoleMouth() {
	m := mouth.NewConsoleMouth()
	m.Say("Hello")
}

func ExampleFileMouth() {
	m := mouth.NewFileMouth("./testfile")
	m.Say("Hello")
}

func main() {
	// ExampleConsoleMouth()
	ExampleFileMouth()
}
