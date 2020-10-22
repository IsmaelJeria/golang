package handlers

import (
	"gorillamux/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Products esructura del handler
type Products struct {
	l *log.Logger
}

//NewProducts func
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//GetProducts ..
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

//AddProduct ..
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle POST Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

//UpdateProduct ..
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to convert id", http.StatusBadRequest)
		return
	}
	p.l.Println("handle PUT Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(rw, "unable to update data", http.StatusBadRequest)
		return
	}
}
