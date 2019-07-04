package rand

import (
	"crypto/rand"
	"encoding/base64"
)

// RememberTokenBytes is a general number large enough
// for our tokens
const RememberTokenBytes = 32

// Bytes will generate n random bytes
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// String will convert a length of bytes into a base64 string
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RememberToken will generate remember token of
// a predetermined byte size
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}
