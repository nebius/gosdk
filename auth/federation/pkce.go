package federation

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
)

type PKCECode string

func GenerateRandomString() (string, error) {
	// https://tools.ietf.org/html/rfc7636#section-4.1:
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"
	var buf [32]byte
	for i := range buf {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		buf[i] = letterBytes[n.Int64()]
	}
	return string(buf[:]), nil
}

func GeneratePKCECode() (PKCECode, error) {
	s, err := GenerateRandomString()
	return PKCECode(s), err
}

func (p PKCECode) Challenge() string {
	b := sha256.Sum256([]byte(p))
	return base64.RawURLEncoding.EncodeToString(b[:])
}

func (p PKCECode) Method() string {
	return "S256"
}

func (p PKCECode) Verifier() string {
	return string(p)
}
