package security

import (
	"math/rand"
	"time"
)

const (
	DefaultIdLength   = 15
	DefaultIdAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func NanoID() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, DefaultIdLength)
	max := len(DefaultIdAlphabet)

	for i := range b {
		n := rnd.Intn(max)
		b[i] = DefaultIdAlphabet[n]
	}

	return string(b)
}
