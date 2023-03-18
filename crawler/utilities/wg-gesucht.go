package utilities

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yzaimoglu/flathunter/models/crawler"
)

func GetDetailsWGGesucht(details []string, listing models.Listing) (resultingListing models.Listing) {
  // Loop through scraped details and harvest specific details 
  for i := range(details) {
    details[i] = strings.ReplaceAll(strings.ReplaceAll(details[i], " ", ""), "\n", "")
    replacer := ""
    if strings.HasPrefix(details[i], "Wohnfläche") {
      replacer = "Wohnfläche"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      listing.Size = details[i]
    } else if strings.HasPrefix(details[i], "Zimmer") {
      replacer = "Zimmer"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      rooms_int, err := strconv.Atoi(details[i])
      if err != nil {
        fmt.Println("Error during conversion of rooms to int")
        rooms_int = 0
      }
      listing.Rooms = rooms_int
    } else if strings.HasPrefix(details[i], "Badezimmer") {
      replacer = "Badezimmer"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      bathrooms_int, err := strconv.Atoi(details[i])
      if err != nil {
        fmt.Println("Error during conversion of bathrooms to int")
        bathrooms_int = 0
      }
      listing.Bathrooms = bathrooms_int
    } else if strings.HasPrefix(details[i], "Etage") {
      replacer = "Etage"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      listing.Floor = details[i]
    } else if strings.HasPrefix(details[i], "Wohnungstyp") {
      replacer = "Wohnungstyp"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      listing.Type = details[i]
    } else if strings.HasPrefix(details[i], "Nebenkosten") {
      replacer = "Nebenkosten"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      listing.ExtraCosts = details[i]
    } else if strings.HasPrefix(details[i], "Warmmiete") {
      replacer = "Warmmiete"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      listing.FullRent = details[i]
    } else if strings.HasPrefix(details[i], "Kaution/Genoss.-Anteile") {
      replacer = "Kaution/Genoss.-Anteile"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      listing.Deposit = details[i]
    } else if strings.HasPrefix(details[i], "Schlafzimmer") {
      replacer = "Schlafzimmer"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      bedrooms_int, err := strconv.Atoi(details[i])
      if err != nil {
        fmt.Println("Error during conversion of bedrooms to int")
        bedrooms_int = 0
      }
      listing.Bedrooms = bedrooms_int
    } else if strings.HasPrefix(details[i], "Verfügbarab") {
      replacer = "Verfügbarab"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      listing.Availability = details[i]
    } else if strings.HasPrefix(details[i], "Online-Besichtigung") {
      replacer = "Online-Besichtigung"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      listing.OnlineTour = details[i]
    } else if strings.HasPrefix(details[i], "Heizkosten") {
      replacer = "Heizkosten"
      details[i] = strings.Replace(details[i], replacer, "", 1)
      listing.HeatingCosts = details[i]
    }
    
  }

  return listing
}
