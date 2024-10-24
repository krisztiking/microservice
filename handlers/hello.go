package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world!")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if len(d) == 0 {
		http.Error(rw, "Request body cannot be empty", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)

}
