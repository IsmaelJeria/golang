package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

//Product definicion de producton es la estructura de la API
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedON   string  `json:"-"`
}

//Products ..
type Products []*Product

//FromJSON ..
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//ToJSON ..
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

//AddProduct ..
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

//UpdateProduct ..
func UpdateProduct(id int, p *Product) error {
	up, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = up.ID
	productList[pos] = p
	return nil
}

//ErrProductNotFound ..
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	fmt.Println(id)
	{
		for i, p := range productList {
			if p.ID == id {
				return p, i, nil
			}
		}
		return nil, -1, ErrProductNotFound
	}
}

//GetProducts obtiene productos
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "expresso",
		Description: "cafe neeeegro",
		Price:       1.99,
		SKU:         "deg456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
