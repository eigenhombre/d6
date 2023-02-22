package main

// Loose translation of "Nominal" Common Lisp repo
// https://github.com/eigenhombre/nominal/blob/master/src/main.lisp

import (
	_ "embed"
	"math/rand"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed names.txt
var Names string

func nameList() []string {
	return strings.Split(Names, "\n")
}

func tradName() string {
	return randNth(nameList())
}

func randNth(list []string) string {
	return list[randSrc().Intn(len(list))]
}

func singleName(r *rand.Rand) string {
	if r.Intn(4) == 0 {
		return nGramName(r)
	}
	return tradName()
}

func stringListNgrams(list []string, n int) []string {
	ngrams := []string{}
	for _, s := range list {
		ngrams = append(ngrams, stringNgrams(s, n)...)
	}
	return ngrams
}

func stringNgrams(s string, n int) []string {
	ngrams := []string{}
	for i := 0; i < len(s)-n+1; i++ {
		ngrams = append(ngrams, s[i:i+n])
	}
	return ngrams
}

func nGramFrequencies(nGrams []string) map[string]int {
	freqs := map[string]int{}
	for _, n := range nGrams {
		freqs[n]++
	}
	return freqs
}

type pair struct {
	key  string
	freq int
}

func sortedFrequencies(freqs map[string]int) []pair {
	pairs := []pair{}
	for k, v := range freqs {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})
	return pairs
}

var topNgramsGlobal []string = []string{}

func topNgrams(nGrams []string, n int) []string {
	freqs := nGramFrequencies(nGrams)
	pairs := sortedFrequencies(freqs)
	top := []string{}
	for i := 0; i < n; i++ {
		if i >= len(pairs) {
			break
		}
		top = append(top, pairs[i].key)
	}
	return top
}

func makeNgrams(corpus []string, n int) []string {
	return topNgrams(stringListNgrams(corpus, n), 200)
}

func allNgrams(list []string) []string {
	ngrams := []string{}
	for n := 3; n < 5; n++ {
		ngrams = append(ngrams, makeNgrams(list, n)...)
	}
	return ngrams
}

func nGramName(r *rand.Rand) string {
	numgrams := r.Intn(r.Intn(4)+1) + 1
	name := ""
	for i := 0; i < numgrams; i++ {
		name += randNth(topNgramsGlobal)
	}
	return name
}

func fullNameAsList(r *rand.Rand) []string {
	ns := []int{1, 2, 2, 2, 2, 2, 3, 3, 4, 5, 6}
	numNames := ns[r.Intn(len(ns))]
	names := []string{}
	c := cases.Title(language.English)
	for i := 0; i < numNames; i++ {
		names = append(names, c.String(singleName(r)))
	}
	return names
}

func calculateTopNgrams() {
	topNgramsGlobal = topNgrams(allNgrams(nameList()), 1000)
}

func fullName(r *rand.Rand) string {
	return strings.Join(fullNameAsList(r), " ")
}
