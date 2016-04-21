package ear

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rikonor/people/person/dto"
)

// HTTPEar - serves as Handler interface between the person and the HTTPWorld
type HTTPEar struct {
	Ear
	Port string
}

// NewHTTPEar - Create a new HTTPEar
func NewHTTPEar(port string) *HTTPEar {
	he := HTTPEar{
		Port: port,
	}

	go he.Listen()

	return &he
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

func (he *HTTPEar) Listen() {
	sMux := http.NewServeMux()
	sMux.HandleFunc("/", getHTTPHandler(he.relayMessage))
	http.ListenAndServe(":"+he.Port, sMux)
}
