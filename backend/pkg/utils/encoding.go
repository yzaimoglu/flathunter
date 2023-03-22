package utils

import "encoding/base64"

// ToBase64 encodes a byte array to a base64 string
func ToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// FromBase64 decodes a base64 string to a byte array
func FromBase64(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
