package models

// Listing represents a listing of a flat.
type Listing struct {
	ID           int64    `json:"id"`
	URLID        int64    `json:"url_id"`
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
	CrawledAt    int64    `json:"crawled_at"`
}

// UserListing represents a user listing.
type UserListing struct {
	ID        int64 `json:"id"`
	UserID    int64 `json:"user_id"`
	ListingID int64 `json:"listing_id"`
	Notified  bool  `json:"notified"`
}
