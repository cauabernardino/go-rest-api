package utils

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFloat generates a random float between a min and max given as integers
func randomFloat(min, max int64) float64 {
	intPart := float64(rand.Int63n(max - min + 1))

	return float64(min) + intPart + rand.Float64()
}

// randomString generates a random string of length n
func randomString(n int) string {
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
	return randomString(10)
}

// RandomPrice generates a random amount of money
func RandomPrice() float64 {
	price := randomFloat(0, 1000)

	return math.Round(price*100) / 100
}

// RandomDescription generates a random description
func RandomDescription() string {
	return randomString(80)
}
