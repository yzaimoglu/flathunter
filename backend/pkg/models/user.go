/*
models encompasses all required struct types.
This package is used by the backend, as well as the crawler to rightfully format JSON Objects
into Go structs. It also has necessary utility classes.
*/
package models

// User is the representation of a User inside of Go and in the Database.
type User struct {
	ID   string `json:"_id"`
	Key  string `json:"_key"`
	Rev  string `json:"_rev"`
	Name string `json:"name"`
	Role Role   `json:"role"`
}
