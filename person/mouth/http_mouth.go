package mouth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rikonor/people/person/dto"
)

type HTTPMouth struct {
	Mouth
	// Host - format <host>:<port>
	// Example - localhost:8080
	Host string
}

func NewHTTPMouth(host string) *HTTPMouth {
	m := HTTPMouth{Host: host}
	m.InputChannel = make(chan dto.Message, DefaultInputBufferCapacity)
	go m.waitForSomethingToSay(m.SaySync)
	return &m
}

func (m *HTTPMouth) SaySync(msg dto.Message) {
	_, err := http.Post("http://"+m.Host, "text/plain", strings.NewReader(string(msg)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send message to HTTP endpoint: http://%s: %s\n", m.Host, err)
	}
}
