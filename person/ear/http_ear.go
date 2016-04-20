package ear

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rikonor/people/person/dto"
)

// HTTPEar - serves as Handler interface between the person and the HTTPWorld
type HTTPEar struct {
	MessageHandler dto.MessageHandler
	Port           string
}

// NewHTTPEar - Create a new HTTPEar
func NewHTTPEar(msgHandler dto.MessageHandler, port string) *HTTPEar {
	he := HTTPEar{
		MessageHandler: msgHandler,
		Port:           port,
	}

	go he.listen()

	return &he
}

func (he *HTTPEar) relayMessage(msg dto.Message) {
	if he.MessageHandler == nil {
		fmt.Println("Ear can't relay message since no message handler was defined")
		return
	}
	he.MessageHandler(msg)
}

func getHTTPHandler(msgHandler dto.MessageHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		msg := dto.ByteSliceToMessage(data)
		msgHandler(msg)
		fmt.Fprint(w, "OK")
	}
}

func (he *HTTPEar) listen() {
	http.HandleFunc("/", getHTTPHandler(he.relayMessage))
	http.ListenAndServe(":"+he.Port, nil)
}
