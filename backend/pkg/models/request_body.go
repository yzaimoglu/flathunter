package models

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateURL struct {
	ArangoModel
	Platform    string `json:"platform"`
	URL         string `json:"url"`
	CreatedAt   int64  `json:"created_at"`
	LastCrawled int64  `json:"last_crawled"`
}
