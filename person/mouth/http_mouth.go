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
	Hosts []string
}

func NewHTTPMouth(hosts ...string) *HTTPMouth {
	m := HTTPMouth{Hosts: hosts}
	m.InputChannel = make(chan dto.Message, DefaultInputBufferCapacity)
	go m.waitForSomethingToSay(m.SaySync)
	return &m
}

func (m *HTTPMouth) SaySync(msg dto.Message) {
	for _, host := range m.Hosts {
		_, err := http.Post("http://"+host, "text/plain", strings.NewReader(string(msg)))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to send message to HTTP endpoint: http://%s: %s\n", host, err)
		}
	}
}
