package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func upp(rand *rand.Rand) string {
	ret := ""
	for i := 0; i < 6; i++ {
		ret += fmt.Sprintf("%X", dn(rand, 2))
	}
	return ret
}
func chars(r *rand.Rand) string {
	ret := []string{}
	for i := 0; i < 10; i++ {
		ret = append(ret, fmt.Sprintf("%s %s", upp(r), fullName(r)))
	}
	return strings.Join(ret, "\n")
}
