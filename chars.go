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
	events    []lifeEvent
}

type skill struct {
	name  string
	level int
}

func professions() []string {
	return []string{
		"aerospace system defense",
		"agent",
		"athlete",
		"barbarian",
		"belter",
		"bureaucrat",
		"colonist",
		"diplomat",
		"drifter",
		"entertainer",
		"hunter",
		"marine",
		"maritime system defense",
		"mercenary",
		"merchant",
		"navy",
		"noble",
		"physician",
		"pirate",
		"rogue",
		"scientist",
		"scout",
		"surface system defense",
		"technician",
	}
}

func newChar(r *rand.Rand) char {
	firstService := randNthString(r, professions())
	return char{
		name:      fullName(r),
		age:       18,
		st:        dn(r, 2),
		dx:        dn(r, 2),
		en:        dn(r, 2),
		in:        dn(r, 2),
		ed:        dn(r, 2),
		ss:        dn(r, 2),
		career:    firstService,
		homeworld: newPlanet(r, nil, nil),
		events: []lifeEvent{
			birthEvent{},
			majorityEvent{},
			attemptToJoinServiceEvent{
				service: firstService,
				age:     18,
			},
		},
		skills: []skill{},
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
	ret := []string{}
	ret = append(ret, fmt.Sprintf("%s %X%X%X%X%X%X, from %s (%s), %d y.o., %s. %s",
		c.name, c.st, c.dx, c.en, c.in, c.ed, c.ss,
		c.homeworld.name, c.homeworld.uwp(),
		c.age, c.career,
		skillsList(c.skills)))
	for _, e := range c.events {
		ret = append(ret, fmt.Sprintf("Age %d: %s", e.Age(), e.Name()))
	}
	return strings.Join(ret, "\n")
}

func chars(r *rand.Rand) string {
	ret := []string{}
	for i := 0; i < 10; i++ {
		ret = append(ret, newChar(r).String())
	}
	return strings.Join(ret, "\n")
}
