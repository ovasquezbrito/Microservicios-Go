package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Hello Hello para saludar
type Hello struct {
	l *log.Logger
}

// NewHello - funcion la el modelo
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Running Hello Handler")

	// read the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)

		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(w, "Hello %s", b)
}
