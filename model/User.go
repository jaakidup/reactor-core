package model

// User ...
type User struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Age     int    `json:"age,omitempty"`
}
