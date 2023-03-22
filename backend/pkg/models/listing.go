package models

// Platform is a struct that holds the information for a platform.
type Platform struct {
	ArangoModel
	Name         string `json:"name"`
	ReadableName string `json:"readable_name"`
}

// URL is a struct that holds the information for a URL.
type URL struct {
	ArangoModel
	Platform    Platform `json:"platform"`
	URL         string   `json:"url"`
	CreatedAt   int64    `json:"created_at"`
	LastCrawled int64    `json:"last_crawled"`
}

// UserURL is a struct that holds the information for a user URL.
type UserURL struct {
	ArangoModel
	User User `json:"user"`
	URL  URL  `json:"url_id"`
}

// Listing is a struct that holds the information for a listing.
type Listing struct {
	ArangoModel
	URL          URL      `json:"url"`
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

// UserListing is a struct that holds the information for a user listing.
type UserListing struct {
	ArangoModel
	User     User    `json:"user"`
	Listing  Listing `json:"listing"`
	Notified bool    `json:"notified"`
}
