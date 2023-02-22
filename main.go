package main

import (
	"fmt"
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

func upp(rand *rand.Rand) string {
	ret := ""
	for i := 0; i < 6; i++ {
		ret += fmt.Sprintf("%X", dn(rand, 2))
	}
	return ret
}

func main() {
	r := randSrc()
	fmt.Println(upp(r))
}
