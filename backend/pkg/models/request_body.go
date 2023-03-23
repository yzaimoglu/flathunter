package models

type CreateUser struct {
	ArangoModel
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

type CreateListing struct {
	ArangoModel
	URL          string   `json:"url"`
	Images       []string `json:"images"`
	Price        string   `json:"price"`
	Date         string   `json:"date"`
	Description  string   `json:"description"`
	Size         string   `json:"size"`
	Rooms        int      `json:"rooms"`
	Bathrooms    int      `json:"bathrooms"`
	Floor        string   `json:"floor"`
	Type         string   `json:"type"`
	ExtraCosts   string   `json:"extra_costs"`
	FullRent     string   `json:"full_rent"`
	Deposit      string   `json:"deposit"`
	Bedrooms     int      `json:"bedrooms"`
	OnlineTour   string   `json:"online_tour"`
	HeatingCosts string   `json:"heating_costs"`
	Availability string   `json:"availability"`
	CreatedAt    int64    `json:"created_at"`
}
