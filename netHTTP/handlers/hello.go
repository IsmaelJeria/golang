package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Hello esructura del handler
type Hello struct {
	l *log.Logger
}

//NewHello func
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("hola mundo")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		//		rw.WriteHeader(http.StatusBadRequest)
		//		rw.Write([]byte("oooops"))
		return
	}
	fmt.Fprintf(rw, "Data %s\n", d)
}
