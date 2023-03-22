package models

// ArangoModel is a struct that contains the basic fields of an ArangoDB document.
type ArangoModel struct {
	ID  string `json:"_id,omitempty"`
	Key string `json:"_key,omitempty"`
	Rev string `json:"_rev,omitempty"`
}
