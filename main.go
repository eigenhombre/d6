package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func upp(rand *rand.Rand) string {
	ret := ""
	for i := 0; i < 6; i++ {
		ret += fmt.Sprintf("%X", dn(rand, 2))
	}
	return ret
}

func usage() {
	fmt.Println("Usage: d6 <verb> [verb] ...")
	fmt.Println("Verbs:")
	fmt.Println("\tchars")
	fmt.Println("\tworlds")
	os.Exit(1)
}

func chars(r *rand.Rand) string {
	ret := []string{}
	for i := 0; i < 10; i++ {
		ret = append(ret, fmt.Sprintf("%s %s", upp(r), fullName(r)))
	}
	return strings.Join(ret, "\n")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type planet struct {
	name     string
	size     int
	atmo     int
	hydro    int
	pop      int
	gov      int
	law      int
	tl       int
	starport rune
}

func newPlanet(r *rand.Rand) planet {
	size := dn(r, 2)
	var atmo int
	if size == 0 {
		atmo = 0
	} else {
		atmo = max(0, min(dn(r, 2)-7+size, 15))
	}
	hydro_dm := 0
	var hydro int
	if size < 2 {
		hydro = 0
	}
	if atmo == 0 || atmo == 1 || atmo == 10 || atmo == 11 || atmo == 12 {
		hydro_dm = -4
	} else if atmo == 14 {
		hydro_dm = -2
	}
	hydro = max(0, min(dn(r, 2)-7+size+hydro_dm, 10))
	pop_dm := 0
	if size < 3 {
		pop_dm = -1
	}
	if atmo > 9 {
		pop_dm += -2
	}
	if atmo == 6 {
		pop_dm += 3
	}
	if atmo == 5 || atmo == 8 {
		pop_dm += 1
	}
	if hydro == 0 && atmo < 3 {
		pop_dm += -2
	}
	pop := max(0, min(dn(r, 2)-7+size+pop_dm, 10))
	gov := max(0, min(dn(r, 2)-7+size, 15))
	law := max(0, min(dn(r, 2)-7+gov, 15))
	sp_num := dn(r, 2) - 7 + pop
	if sp_num < 3 {
		sp_num = 2
	}
	if sp_num > 11 {
		sp_num = 11
	}
	sp_map := map[int]rune{
		2:  'X',
		3:  'E',
		4:  'E',
		5:  'D',
		6:  'D',
		7:  'C',
		8:  'C',
		9:  'B',
		10: 'B',
		11: 'A',
	}
	starport := sp_map[sp_num]
	tl_dm := 0
	if starport == 'A' {
		tl_dm += 6
	} else if starport == 'B' {
		tl_dm += 4
	} else if starport == 'C' {
		tl_dm += 2
	} else if starport == 'X' {
		tl_dm -= 4
	}
	if size == 0 || size == 1 {
		tl_dm += 2
	}
	if size == 2 || size == 3 || size == 4 {
		tl_dm += 1
	}
	if atmo < 4 || atmo > 9 {
		tl_dm += 1
	}
	if hydro == 0 {
		tl_dm += 1
	}
	if hydro == 9 {
		tl_dm += 1
	}
	if hydro == 10 {
		tl_dm += 2
	}
	if pop > 0 && pop < 6 {
		tl_dm += 1
	}
	if pop == 9 {
		tl_dm += 1
	}
	if pop == 10 {
		tl_dm += 2
	}
	if pop == 11 {
		tl_dm += 3
	}
	if pop == 12 {
		tl_dm += 4
	}
	if gov == 0 || gov == 5 {
		tl_dm += 1
	}
	if gov == 13 || gov == 14 {
		tl_dm -= 2
	}
	tl := max(0, min(15, d(r)+tl_dm))
	// Check minimum TL
	if (hydro == 0 || hydro == 10) && pop > 5 {
		tl = max(tl, 4)
	}
	if atmo == 4 || atmo == 7 || atmo == 9 {
		tl = max(tl, 5)
	}
	if atmo < 4 || atmo > 9 {
		tl = max(tl, 7)
	}
	if (atmo == 13 || atmo == 14) && hydro == 10 {
		tl = max(tl, 7)
	}
	return planet{
		name:     capitalize(singleName(r)),
		size:     size,
		atmo:     atmo,
		hydro:    hydro,
		pop:      pop,
		gov:      gov,
		law:      law,
		tl:       tl,
		starport: starport,
	}
}

func (p planet) String() string {
	return fmt.Sprintf("%20s    0000 %c%X%X%X%X%X%X-%X",
		p.name,
		p.starport,
		p.size,
		p.atmo,
		p.hydro,
		p.pop,
		p.gov,
		p.law,
		p.tl,
	)
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
