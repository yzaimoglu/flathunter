package models

// Proxy is the representation of a Proxy inside of Go and in the Database.
type Proxy struct {
	ArangoModel
	IP       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

// UserAgent is the representation of a UserAgent inside of Go and in the Database.
type UserAgent struct {
	ArangoModel
	UserAgent string `json:"string"`
}
