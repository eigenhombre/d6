package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type planet struct {
	name       string
	sectorX    int
	sectorY    int
	size       int
	atmo       int
	hydro      int
	pop        int
	gov        int
	law        int
	tl         int
	starport   rune
	tradeCodes []string
}

func newPlanet(r *rand.Rand, sectorX, sectorY int) planet {
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
	var gov int
	if pop > 0 {
		gov = max(0, min(dn(r, 2)-7+size, 15))
	}
	var law int
	if pop > 0 {
		law = max(0, min(dn(r, 2)-7+gov, 15))
	}
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
	var tl int
	if pop == 0 {
		tl = 0
	} else {
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
		tl = max(0, min(15, d(r)+tl_dm))
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
	}
	tradeCodes := calculateTradeCodes(
		size,
		atmo,
		hydro,
		pop,
		gov,
		law,
		tl,
	)
	return planet{
		name:       capitalize(singleName(r)),
		sectorX:    sectorX,
		sectorY:    sectorY,
		size:       size,
		atmo:       atmo,
		hydro:      hydro,
		pop:        pop,
		gov:        gov,
		law:        law,
		tl:         tl,
		starport:   starport,
		tradeCodes: tradeCodes,
	}
}

func calculateTradeCodes(size, atmo, hydro, pop, gov, law, tl int) []string {
	ret := []string{}
	if 4 <= atmo && atmo <= 9 && 4 <= hydro && hydro <= 8 && 5 <= pop && pop <= 7 {
		ret = append(ret, "Ag")
	}
	if size == 0 && atmo == 0 && hydro == 0 {
		ret = append(ret, "As")
	}
	if pop == 0 && gov == 0 && law == 0 {
		ret = append(ret, "Ba")
	}
	if atmo >= 2 && hydro == 0 {
		ret = append(ret, "De")
	}
	if atmo >= 10 && hydro >= 1 {
		ret = append(ret, "Fl")
	}
	if (atmo == 5 || atmo == 6 || atmo == 8) && hydro > 3 && hydro < 10 && pop > 3 && pop < 9 {
		ret = append(ret, "Ga")
	}
	if pop >= 9 {
		ret = append(ret, "Hi")
	}
	if tl >= 12 {
		ret = append(ret, "Ht")
	}
	if atmo < 3 && hydro >= 1 {
		ret = append(ret, "Ic")
	}
	if (atmo < 3 || atmo == 4 || atmo == 7 || atmo == 9) && pop >= 9 {
		ret = append(ret, "In")
	}
	if pop >= 1 && pop <= 3 {
		ret = append(ret, "Lo")
	}
	if tl <= 5 && pop > 0 {
		ret = append(ret, "Lt")
	}
	if atmo <= 3 && hydro <= 3 && pop >= 6 {
		ret = append(ret, "Na")
	}
	if pop >= 4 && pop <= 6 {
		ret = append(ret, "Ni")
	}
	if atmo >= 2 && atmo <= 5 && hydro <= 3 && pop > 0 {
		ret = append(ret, "Po")
	}
	if (atmo == 6 || atmo == 8) && pop >= 6 && pop <= 8 {
		ret = append(ret, "Ri")
	}
	if hydro == 10 {
		ret = append(ret, "Wa")
	}
	if atmo == 0 {
		ret = append(ret, "Va")
	}
	return ret
}

func (p planet) String() string {
	return fmt.Sprintf("%20s    %02d%02d %c%X%X%X%X%X%X-%X %s",
		p.name,
		p.sectorX,
		p.sectorY,
		p.starport,
		p.size,
		p.atmo,
		p.hydro,
		p.pop,
		p.gov,
		p.law,
		p.tl,
		strings.Join(p.tradeCodes, " "),
	)
}

func generateSubsector(r *rand.Rand) string {
	ret := []string{}
	for sectorX := 0; sectorX < 8; sectorX++ {
		for sectorY := 0; sectorY < 10; sectorY++ {
			if r.Intn(2) == 0 {
				ret = append(ret, newPlanet(r, sectorX, sectorY).String())
			}
		}
	}
	return strings.Join(ret, "\n")
}
