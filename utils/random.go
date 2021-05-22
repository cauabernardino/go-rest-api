package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomFloat generates a random float between a min and max given as integers
func RandomFloat(min, max int64) float64 {
	intPart := float64(rand.Int63n(max - min + 1))

	return float64(min) + intPart + rand.Float64()
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName generates a random owner name
func RandomName() string {
	return RandomString(10)
}

// RandomPrice generates a random amount of money
func RandomPrice() float64 {
	return RandomFloat(0, 1000)
}

// RandomDescription generates a random description
func RandomDescription() string {
	return RandomString(80)
}
