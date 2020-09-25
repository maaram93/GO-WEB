package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello struct implements ServeHttp method which makes it a handler.
type Hello struct {
	l *log.Logger
}

// NewHello creates a new Hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP method is the implementation which handles request and response.
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello Maaram")

	// read body from request
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops Badrequest", http.StatusBadRequest)
		return
	}
	h.l.Printf("Data %s\n", d)

	// write the response
	fmt.Fprintf(rw, "Hello %s", d)
}
