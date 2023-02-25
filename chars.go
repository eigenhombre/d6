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

var characteristicMods map[int]int = map[int]int{
	0:  -2,
	1:  -2,
	2:  -1,
	3:  -1,
	4:  -1,
	5:  -1,
	6:  0,
	7:  0,
	8:  0,
	9:  1,
	10: 1,
	11: 1,
	12: 2,
	13: 2,
	14: 2,
	15: 3,
	16: 3,
	17: 3,
	18: 4,
	19: 4,
	20: 4,
}

type skill struct {
	name  string
	level int
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
		careers:   []career{},
		homeworld: newPlanet(r, nil, nil),
		events: []lifeEvent{
			birthEvent{},
			majorityEvent{},
		},
		skills: []skill{},
	}
}

func charCareer(r *rand.Rand, c *char) {
	service := randCareer(r, allCareers())
	c.careers = append(c.careers, service)
	if !enlist(r, service, *c) {
		c.events = append(c.events, failedToJoinServiceEvent{
			service: service,
			age:     c.age,
		})
		if d(r) < 4 {
			c.events = append(c.events, draftedEvent{
				service: drifter,
				age:     c.age,
			})
			c.careers = append(c.careers, drifter)
		} else {
			service := randCareer(r, draftCareers)
			c.events = append(c.events, draftedEvent{
				service: service,
				age:     c.age,
			})
			c.careers = append(c.careers, service)
		}
		return
	}
	c.events = append(c.events, joinedServiceEvent{
		service: service,
		age:     c.age,
	})
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
		ret = append(ret, fmt.Sprintf("%3d - %s", e.Age(), e.Name()))
	}
	return strings.Join(ret, "\n") + "\n"
}

func chars(r *rand.Rand) string {
	ret := []string{}
	for i := 0; i < 10; i++ {
		ch := newChar(r)
		charCareer(r, &ch)
		ret = append(ret, ch.String())
	}
	return strings.Join(ret, "\n")
}
