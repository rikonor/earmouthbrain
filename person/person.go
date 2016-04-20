package person

import (
	"fmt"
	"net/http"
	"strconv"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

// New creates a new Person
func New(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) String() string {
	return fmt.Sprintf("%s", p.Name)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Only GET accepted")
	}
}

func (p *Person) Listen(port int) {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				fmt.Fprintf(w, "Hello")
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "Error: Only GET accepted")
			}
		})

		p.Address = "http://localhost:" + strconv.Itoa(port)

		http.ListenAndServe(":"+strconv.Itoa(port), nil)
	}()
}

func (pFrom *Person) TalkTo(pTo *Person) {
	res, err := http.Get(pTo.Address)
	if err != nil {

	}

	if res.StatusCode != http.StatusOK {

	}
}
