package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type char struct {
	name string
	age  int
	st   int
	dx   int
	en   int
	in   int
	ed   int
	ss   int
}

func newChar(r *rand.Rand) char {
	return char{
		name: fullName(r),
		age:  18,
		st:   dn(r, 2),
		dx:   dn(r, 2),
		en:   dn(r, 2),
		in:   dn(r, 2),
		ed:   dn(r, 2),
		ss:   dn(r, 2),
	}
}

func (c char) String() string {
	return fmt.Sprintf("%s %X%X%X%X%X%X, %d y.o.", c.name, c.st, c.dx, c.en, c.in, c.ed, c.ss, c.age)
}

func chars(r *rand.Rand) string {
	ret := []string{}
	for i := 0; i < 10; i++ {
		ret = append(ret, newChar(r).String())
	}
	return strings.Join(ret, "\n")
}
