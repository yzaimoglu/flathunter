package models

type CreateUser struct {
	ArangoModel
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type CreateURL struct {
	ArangoModel
	Platform    string `json:"platform,omitempty"`
	URL         string `json:"url,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	LastCrawled int64  `json:"last_crawled,omitempty"`
}

type CreateListing struct {
	ArangoModel
	URL          string   `json:"url,omitempty"`
	Images       []string `json:"images,omitempty"`
	Price        string   `json:"price,omitempty"`
	Date         string   `json:"date,omitempty"`
	Description  string   `json:"description,omitempty"`
	Size         string   `json:"size,omitempty"`
	Rooms        int      `json:"rooms,omitempty"`
	Bathrooms    int      `json:"bathrooms,omitempty"`
	Floor        string   `json:"floor,omitempty"`
	Type         string   `json:"type,omitempty"`
	ExtraCosts   string   `json:"extra_costs,omitempty"`
	FullRent     string   `json:"full_rent,omitempty"`
	Deposit      string   `json:"deposit,omitempty"`
	Bedrooms     int      `json:"bedrooms,omitempty"`
	OnlineTour   string   `json:"online_tour,omitempty"`
	HeatingCosts string   `json:"heating_costs,omitempty"`
	Availability string   `json:"availability,omitempty"`
	CreatedAt    int64    `json:"created_at,omitempty"`
}
