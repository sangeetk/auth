package token

import (
	"fmt"
	"math/rand"
)

// RandomToken - Generates a random token
func RandomToken(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
