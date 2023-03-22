package models

// Role is the representation of a user role inside of Go and in the Database.
// It has all the necessary information for making protected actions.
type Role struct {
	ID   string `json:"_id"`
	Key  string `json:"_key"`
	Rev  string `json:"_rev"`
	Name string `json:"name"`
}
