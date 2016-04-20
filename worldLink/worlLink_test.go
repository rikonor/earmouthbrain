package worldLink

import (
	"net/http"
	"testing"
)

func TestWorldLinkConnectAndDisconnect(t *testing.T) {
	wl := New()

	wl.Connect()
	defer wl.Disconnect()
}

func TestWorldLinkConnectPopulatesPortField(t *testing.T) {
	wl := New()

	wl.Connect()
	defer wl.Disconnect()

	if wl.Port == 0 {
		t.Error("Failed to bind to a port")
	}
}

func TestWorldLinkCanRespondToIncomingRequests(t *testing.T) {
	wl := New()

	wl.Connect()
	defer wl.Disconnect()

	res, err := http.Get("http://localhost:" + string(wl.Port))
	if err != nil {
		t.Error("Failed to make a call to WorldLink with error:", err)
	}

	if res.StatusCode != 200 {
		t.Error("Failed to make a successful call to WorldLink")
	}
}
