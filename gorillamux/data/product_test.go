package data

import (
	"testing"
)

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "ismael",
		Price: 1.0,
		SKU:   "asd-qwe-asddsadsa",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
