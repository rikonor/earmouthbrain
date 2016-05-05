package earmouthbrain

import (
	"fmt"
	"testing"
)

func testMessageHandler(msg Message) {
	fmt.Println("Captured message")
}

func TestCanRegisterHandler(t *testing.T) {
	e := &Ear{}

	if len(e.MessageHandlers) != 0 {
		t.Error("Failed to confirm fresh instance of Ear has no message handlers")
	}

	e.RegisterMessageHandler(testMessageHandler)

	if len(e.MessageHandlers) != 1 {
		t.Error("Failed to register message handler with ear")
	}
}

func TestCanRelayMessage(t *testing.T) {
	e := &Ear{}

	msgCaptured := false

	e.RegisterMessageHandler(func(_ Message) {
		msgCaptured = true
	})

	msg := StringToMessage("Test")
	e.RelayMessage(msg)

	if !msgCaptured {
		t.Error("Failed to relay message to registered message handlers")
	}
}

func ExampleEar() {
	e := &Ear{}
	e.RegisterMessageHandler(testMessageHandler)

	msg := StringToMessage("Test")
	e.RelayMessage(msg)

	// Output: Captured message
}
