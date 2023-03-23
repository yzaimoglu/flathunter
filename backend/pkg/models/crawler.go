package models

// Proxy is the representation of a Proxy inside of Go and in the Database.
type Proxy struct {
	ArangoModel
	IP       string `json:"ip,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Port     int    `json:"port,omitempty"`
}

// UserAgent is the representation of a UserAgent inside of Go and in the Database.
type UserAgent struct {
	ArangoModel
	UserAgent string `json:"string,omitempty"`
}
