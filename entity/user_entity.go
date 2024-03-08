package entity

import (
	"encoding/json"
	"io"
)

type User struct {
	ID            string `json:"id"`
	Fullname      string `json:"fullname"`
	TaxIdentifier string `json:"tax_identifier"`
	Email         string `json:"email"`
	Password      string `json:"password"`
}

func NewUser(r io.Reader) *User {
	var user_request User
	json.NewDecoder(r).Decode(&user_request)
	return &user_request
}
