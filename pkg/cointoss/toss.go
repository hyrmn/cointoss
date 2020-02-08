package cointoss

import (
	"math/rand"
	"time"
)

// Toss will return either 0 or 1, randomly.
func Toss() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2)
}
