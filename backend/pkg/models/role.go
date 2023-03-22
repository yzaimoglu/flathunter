package models

// Role is the representation of a user role inside of Go and in the Database.
// It has all the necessary information for making protected actions.
type Role struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	ReadableName string `json:"readable_name"`
	Permissions  int64  `json:"permissions"`
}
