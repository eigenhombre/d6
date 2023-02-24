package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
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
		"chars": chars,
		"worlds": func(r *rand.Rand) string {
			ret := []string{}
			for i := 0; i < 10; i++ {
				ret = append(ret, newPlanet(r).String())
			}
			return strings.Join(ret, "\n")
		},
	}
	for v := range verbs {
		_, ok := verbMap[verbs[v]]
		if !ok {
			usage()
		}
		fmt.Println(verbMap[verbs[v]](r))
	}
}
