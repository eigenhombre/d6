package main

import (
	"math/rand"
	"time"
)

func randSrc() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func d(rand *rand.Rand) int {
	return rand.Intn(6) + 1
}

func dn(rand *rand.Rand, n int) int {
	r := 0
	for i := 0; i < n; i++ {
		r += d(rand)
	}
	return r
}
