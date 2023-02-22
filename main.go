package main

import (
	"fmt"
	"math/rand"
)

func upp(rand *rand.Rand) string {
	ret := ""
	for i := 0; i < 6; i++ {
		ret += fmt.Sprintf("%X", dn(rand, 2))
	}
	return ret
}

func main() {
	r := randSrc()
	calculateTopNgrams()
	for i := 0; i < 10; i++ {
		fmt.Println(upp(r) + " " + fullName(r))
	}
}
