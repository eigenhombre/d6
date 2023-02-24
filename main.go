package main

import (
	"fmt"
	"math/rand"
	"os"
)

func usage() {
	fmt.Println("Usage: d6 <verb> [verb] ...")
	fmt.Println("Verbs:")
	fmt.Println("\tchars")
	fmt.Println("\tworlds")
	os.Exit(1)
}

func main() {
	calculateTopNgrams()
	r := randSrc()
	verbs := os.Args[1:]
	if len(verbs) == 0 {
		usage()
	}
	verbMap := map[string]func(rand *rand.Rand) string{
		"chars":  chars,
		"worlds": generateSubsector,
	}
	for v := range verbs {
		_, ok := verbMap[verbs[v]]
		if !ok {
			usage()
		}
		fmt.Println(verbMap[verbs[v]](r))
	}
}
