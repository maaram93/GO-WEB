package handlers

import (
	"log"
	"net/http"
)

// GoodBye struct is and Handler which implements ServeHTTP method.
type GoodBye struct {
	l *log.Logger
}

// NewGoodBye will return a new Goodbye struct
func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeeee"))
}
