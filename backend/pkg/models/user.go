/*
models encompasses all required struct types.
This package is used by the backend, as well as the crawler to rightfully format JSON Objects
into Go structs. It also has necessary utility classes.
*/
package models

// User is the representation of a User inside of Go and in the Database.
type User struct {
	ID             int64  `json:"id"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	CreatedAt      int64  `json:"created_at"`
	LastSeen       int64  `json:"last_seen"`
	TOTP           string `json:"totp"`
	Role           int64  `json:"role"`
}

// UserInformation is the representation of a User inside of Go and in the Database.
// All sensitive information are not included in this type.
type UserInformation struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt int64  `json:"created_at"`
	LastSeen  int64  `json:"last_seen"`
	TOTP      string `json:"totp"`
	Role      int64  `json:"role"`
}

// UserRegistration is the required information when creating a User.
// This data structure is later on used for the User struct.
type UserRegistration struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeated_password"`
}
