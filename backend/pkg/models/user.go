package models

// User is the representation of a User inside of Go and in the Database.
type User struct {
	ArangoModel
	Email          string `json:"email,omitempty"`
	HashedPassword string `json:"hashed_password,omitempty"`
	Role           Role   `json:"role,omitempty"`
}

// CreateUser is the representation of a User used for creating a new User inside of Go and in the Database.
type CreateUser struct {
	ArangoModel
	Email          string `json:"email,omitempty"`
	HashedPassword string `json:"hashed_password,omitempty"`
	Role           string `json:"role,omitempty"`
}

// UserWithSession is the representation of a User with a session token inside of Go and in the Database.
type UserWithSession struct {
	ArangoModel
	User
	SessionToken SessionToken `json:"session_token,omitempty"`
}

// Role is the representation of a user role inside of Go and in the Database.
// It has all the necessary information for making protected actions.
type Role struct {
	ArangoModel
	Name        string `json:"name,omitempty"`
	Permissions int    `json:"permissions,omitempty"`
}
