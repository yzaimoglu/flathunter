package models

type Listing struct {
  ID int64 `json:"id"`
  PlatformID int64 `json:"platform_id"`
  URL string `json:"url"`
  Images []string `json:"images"`
  Price string `json:"price"`
  Date string `json:"date"`
  Description string `json:"description"`
  Size string `json:"size"`
  Rooms int `json:"rooms"`
  Bathrooms int `json:"bathrooms"`
  Floor string `json:"floor"`
  Type string `json:"type"`
  ExtraCosts string `json:"extra_costs"`
  FullRent string `json:"full_rent"`
  Deposit string `json:"deposit"`
  Bedrooms int `json:"bedrooms"`
  OnlineTour string `json:"online_tour"`
  HeatingCosts string `json:"heating_costs"`
  Availability string `json:"availability"`
}
