package models

// SessionToken is the representation of a session token inside of Go and in the Database.
type SessionToken struct {
	ArangoModel
	User         User   `json:"user,omitempty"`
	SessionToken string `json:"session_token,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty"`
	ExpiresAt    int64  `json:"expires_at,omitempty"`
}

// CreateSessionToken is the representation of a session token used for creating a new session token inside of Go and in the Database.
type CreateSessionToken struct {
	ArangoModel
	User         string `json:"user,omitempty"`
	SessionToken string `json:"session_token,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty"`
	ExpiresAt    int64  `json:"expires_at,omitempty"`
}
