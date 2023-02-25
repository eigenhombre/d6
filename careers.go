package main

import "math/rand"

type qual struct {
	trait string
	value int
}

type career struct {
	name          string
	qualification qual
}

func (c career) String() string {
	return c.name
}

var drifter career = career{"drifter", qual{"de", 5}}

var draftCareers []career = []career{
	{"aerospace system defense", qual{"en", 5}},
	{"marine", qual{"in", 6}},
	{"maritime system defense", qual{"en", 5}},
	{"navy", qual{"in", 6}},
	{"scout", qual{"in", 6}},
	{"surface system defense", qual{"en", 5}},
}

var electiveCareers []career = []career{
	{"agent", qual{"ss", 5}},
	{"athlete", qual{"en", 8}},
	{"barbarian", qual{"en", 5}},
	{"belter", qual{"in", 4}},
	{"bureaucrat", qual{"ss", 6}},
	{"colonist", qual{"en", 5}},
	{"diplomat", qual{"ss", 6}},
	drifter,
	{"entertainer", qual{"ss", 8}},
	{"hunter", qual{"en", 5}},
	{"mercenary", qual{"in", 4}},
	{"merchant", qual{"in", 4}},
	{"noble", qual{"ss", 8}},
	{"physician", qual{"ed", 6}},
	{"pirate", qual{"de", 5}},
	{"rogue", qual{"de", 5}},
	{"scientist", qual{"ed", 6}},
	{"technician", qual{"ed", 6}},
}

func allCareers() []career {
	return append(draftCareers, electiveCareers...)
}

func enlist(r *rand.Rand, c career, ch char) bool {
	qual := c.qualification
	traitLookup := map[string]int{
		"st": ch.st,
		"de": ch.dx,
		"en": ch.en,
		"in": ch.in,
		"ed": ch.ed,
		"ss": ch.ss,
	}
	relevant_trait := traitLookup[qual.trait]
	dm := characteristicMods[relevant_trait]
	return dn(r, 2)+dm >= qual.value
}
