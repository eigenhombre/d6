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
	careers   []career
	skills    []skill
	homeworld planet
	events    []lifeEvent
}

type skill struct {
	name  string
	level int
}

func newChar(r *rand.Rand) char {
	firstService := randCareer(r)
	return char{
		name:      fullName(r),
		age:       18,
		st:        dn(r, 2),
		dx:        dn(r, 2),
		en:        dn(r, 2),
		in:        dn(r, 2),
		ed:        dn(r, 2),
		ss:        dn(r, 2),
		careers:   []career{firstService},
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
		c.age, c.careers[0].name,
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
