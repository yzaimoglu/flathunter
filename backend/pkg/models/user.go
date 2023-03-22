package models

// User is the representation of a User inside of Go and in the Database.
type User struct {
	ArangoModel
	Email string `json:"email,omitempty"`
	Role  Role   `json:"role,omitempty"`
}

// Role is the representation of a user role inside of Go and in the Database.
// It has all the necessary information for making protected actions.
type Role struct {
	ID   string `json:"_id"`
	Key  string `json:"_key"`
	Rev  string `json:"_rev"`
	Name string `json:"name"`
}
