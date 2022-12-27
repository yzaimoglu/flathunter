package utilities

import (
	"encoding/base64"
	"log"
)

// Encode bytes to Base64 string
func ToBase64(input []byte) string {
	base64 := base64.URLEncoding.EncodeToString(input)
	return base64
}

// Decode a Base64 string to bytes
func FromBase64(input string) []byte {
	stringBytes, err := base64.URLEncoding.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	return stringBytes
}
