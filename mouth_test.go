package earmouthbrain

import (
	"fmt"
	"testing"
	"time"
)

func TestMouthReadyToProcess(t *testing.T) {
	m := &Mouth{}
	m.Init(func(_ Message) {})

	if !m.isWaitingForSomethingToSay {
		t.Error("Mouth failed to become ready to speak")
	}
}

func ExampleMouth() {
	m := &Mouth{}
	m.Init(func(msg Message) {
		fmt.Println(msg)
	})

	msg := StringToMessage("Test")
	m.Say(msg)

	// Give mouth time to process message
	time.Sleep(1 * time.Millisecond)

	// Output: Test
}
