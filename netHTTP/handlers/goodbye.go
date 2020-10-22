package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//GoodBye esructura del handler
type GoodBye struct {
	l *log.Logger
}

//NewGoodBye func
func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (h *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("adios mundo")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		//		rw.WriteHeader(http.StatusBadRequest)
		//		rw.Write([]byte("oooops"))
		return
	}
	fmt.Fprintf(rw, "adios %s\n", d)
}
