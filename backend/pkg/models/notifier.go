package models

// Notifier is the representation of a Notifier inside of Go and in the Database.
type Notifier struct {
	ArangoModel
	Type    string      `json:"type,omitempty"`
	Options interface{} `json:"options,omitempty"`
	User    User        `json:"user,omitempty"`
}

// CreateNotifier is the representation of a Notifier used for creating a new Notifier inside of Go and in the Database.
type CreateNotifier struct {
	ArangoModel
	Type    string      `json:"type,omitempty"`
	Options interface{} `json:"options,omitempty"`
	User    string      `json:"user,omitempty"`
}
