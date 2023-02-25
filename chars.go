package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type char struct {
	name      string
	age       int
	st        int
	dx        int
	en        int
	in        int
	ed        int
	ss        int
	career    string
	skills    []skill
	homeworld planet
}

type skill struct {
	name  string
	level int
}

func professions() []string {
	return []string{
		"Aerospace System Defense",
		"Agent",
		"Athlete",
		"Barbarian",
		"Belter",
		"Bureaucrat",
		"Colonist",
		"Diplomat",
		"Drifter",
		"Entertainer",
		"Hunter",
		"Marine",
		"Maritime System Defense",
		"Mercenary",
		"Merchant",
		"Navy",
		"Noble",
		"Physician",
		"Pirate",
		"Rogue",
		"Scientist",
		"Scout",
		"Surface System Defense",
		"Technician",
	}
}

func newChar(r *rand.Rand) char {
	return char{
		name:      fullName(r),
		age:       18,
		st:        dn(r, 2),
		dx:        dn(r, 2),
		en:        dn(r, 2),
		in:        dn(r, 2),
		ed:        dn(r, 2),
		ss:        dn(r, 2),
		career:    randNthString(r, professions()),
		homeworld: newPlanet(r, nil, nil),
		skills:    []skill{},
	}
}

func skillsList(skills []skill) string {
	ret := []string{}
	for _, s := range skills {
		ret = append(ret, fmt.Sprintf("%s-%d", s.name, s.level))
	}
	return strings.Join(ret, ", ")
}

func (c char) String() string {

	return fmt.Sprintf("%s %X%X%X%X%X%X, from %s (%s), %d y.o., %s. %s",
		c.name, c.st, c.dx, c.en, c.in, c.ed, c.ss,
		c.homeworld.name, c.homeworld.uwp(),
		c.age, c.career,
		skillsList(c.skills))
}

func chars(r *rand.Rand) string {
	ret := []string{}
	for i := 0; i < 10; i++ {
		ret = append(ret, newChar(r).String())
	}
	return strings.Join(ret, "\n")
}
