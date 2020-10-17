package helpers

import "encoding/base64"

// EncodeBase64 - encode bytes to base64 string
func EncodeBase64(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

// DecodeBase64 - decode base64 string
func DecodeBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

// EncodeURLBase64 - encode bytes to base64 string
func EncodeURLBase64(str []byte) string {
	return base64.URLEncoding.EncodeToString(str)
}

// DecodeURLBase64 - decode base64 string
func DecodeURLBase64(str string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(str)
}
