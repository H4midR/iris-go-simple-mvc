package hfuncs

import (
	"crypto/rand"
	"fmt"
)

// helper functions
//Tokengenerator() : 32byte token generator
func Tokengenerator() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

//CTokengenerator(n int ) : n byte token generator
func CTokengenerator(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
