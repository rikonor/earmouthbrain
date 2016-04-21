package main

import (
	"fmt"

	"github.com/rikonor/people/person/body"
	"github.com/rikonor/people/person/ear"
	"github.com/rikonor/people/person/mouth"
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

func Example5() {
	body.NewBrainlessFileToHTTPBody("./input", "localhost:8080")

	var x string
	fmt.Scanln(&x)
}

func ExampleComplex() {
	eIn := ear.NewFileEar("./input")

	m1 := mouth.NewHTTPMouth("localhost:8080", "localhost:8081")
	eIn.RegisterMessageHandler(m1.SayAsync)

	e1 := ear.NewHTTPEar("8080")
	mOut1 := mouth.NewFileMouth("./output1")
	e1.RegisterMessageHandler(mOut1.SayAsync)

	e2 := ear.NewHTTPEar("8081")
	mOut2 := mouth.NewFileMouth("./output2")
	e2.RegisterMessageHandler(mOut2.SayAsync)

	fmt.Scanln()
}

func main() {
	// Example1()
	// Example2()
	// Example3()
	// Example4()
	// Example5()
	ExampleComplex()
}
