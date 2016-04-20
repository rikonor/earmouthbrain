package person

import (
	"net/http"
	"testing"
)

func TestNewPerson(t *testing.T) {
	p := New("Or Rikon", 27)

	if p.Name != "Or Rikon" {
		t.Error("Failed to create a Person with the correct name")
	}

	if p.Age != 27 {
		t.Error("Failed to create a Person with the correct age")
	}
}

func TestPersonCanListen(t *testing.T) {
	p := New("Or Rikon", 27)
	p.Listen(8080)

	res, err := http.Get("http://localhost:8080")

	if err != nil {
		t.Error("Failed with error:", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error("Failed with non OK response")
	}
}

func TestTwoPersonsCanCommunicate(t *testing.T) {

}

func testPersonCanFindOtherPeople(t *testing.T) {

}

func testPersonCanExchangeWelcomeMessageWithOtherPerson(t *testing.T) {

}
