package handlers

import (
	"log"
	"net/http"
)

//Goodbye - Adios
type Goodbye struct {
	l *log.Logger
}

//NewGoodbye - Nuevo adios
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Byeeee"))
}
