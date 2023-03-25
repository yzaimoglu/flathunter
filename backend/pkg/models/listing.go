package models

// Platform is a struct that holds the information for a platform.
type Platform struct {
	ArangoModel
	Name         string `json:"name,omitempty"`
	ReadableName string `json:"readable_name,omitempty"`
}

// URL is a struct that holds the information for a URL.
type URL struct {
	ArangoModel
	Platform    Platform `json:"platform,omitempty"`
	URL         string   `json:"url,omitempty"`
	CreatedAt   int64    `json:"created_at,omitempty"`
	LastCrawled int64    `json:"last_crawled,omitempty"`
}

// CreateURL is a struct that holds the information for a URL.
type CreateURL struct {
	ArangoModel
	Platform    string `json:"platform,omitempty"`
	URL         string `json:"url,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	LastCrawled int64  `json:"last_crawled,omitempty"`
}

// UserURL is a struct that holds the information for a user URL.
type UserURL struct {
	ArangoModel
	User User `json:"user,omitempty"`
	URL  URL  `json:"url,omitempty"`
}

// CreateUserURL is a struct that holds the information for a user URL.
type CreateUserURL struct {
	ArangoModel
	User string `json:"user,omitempty"`
	URL  string `json:"url,omitempty"`
}

// Listing is a struct that holds the information for a listing.
type Listing struct {
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

// CreateListing is a struct that holds the information to create a listing.
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

// UserListing is a struct that holds the information for a user listing.
type UserListing struct {
	ArangoModel
	User     User    `json:"user,omitempty"`
	Listing  Listing `json:"listing,omitempty"`
	Notified bool    `json:"notified,omitempty"`
}

// CreateUserListing is a struct that holds the information for a user listing.
type CreateUserListing struct {
	ArangoModel
	User     string `json:"user,omitempty"`
	Listing  string `json:"listing,omitempty"`
	Notified bool   `json:"notified,omitempty"`
}
