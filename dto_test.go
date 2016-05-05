package earmouthbrain

import (
	"fmt"
	"testing"
)

func TestConvertStringToMessage(t *testing.T) {
	msg := StringToMessage("Test")

	if string(msg) != "Test" {
		t.Error("Failed to convert string to message and back")
	}
}

func ExampleStringToMessage() {
	msg := StringToMessage("Test")
	fmt.Println(msg)

	// Output: Test
}

func ExampleByteSliceToMessage() {
	bs := []byte("Test")
	msg := ByteSliceToMessage(bs)
	fmt.Println(msg)

	// Output: Test
}
