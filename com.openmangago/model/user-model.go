package model

import (
	"encoding/json"
	"io"
	"time"
)

//User model..
type User struct {
	Name      string    `firestore:"name,omitempty" json:"name"`
	Lastname  string    `firestore:"lastname,omitempty" json:"lastname"`
	Nickname  string    `firestore:"nickname,omitempty" json:"nickname"`
	Password  string    `firestore:"password,omitempty" json:"password"`
	Birthday  int       `firestore:"birthday,omitempty" json:"birthday"`
	CreatedOn time.Time `firestore:"createdon,omitempty" json:"-"`
	UpdateOn  time.Time `firestore:"updateon,omitempty" json:"-"`
}

//Users ..
type Users []User

//FromJSON ..
func (p *Users) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//ToJSON ..
func (p Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
