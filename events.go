package main

import "fmt"

type lifeEvent interface {
	Name() string
	Age() int
}

type birthEvent struct{}

func (e birthEvent) Name() string {
	return "born"
}

func (e birthEvent) Age() int {
	return 0
}

type majorityEvent struct{}

func (e majorityEvent) Name() string {
	return "majority reached"
}

func (e majorityEvent) Age() int {
	return 18
}

type attemptToJoinServiceEvent struct {
	service career
	age     int
}

func (e attemptToJoinServiceEvent) Name() string {
	return fmt.Sprintf("attempted to begin career as a(n) %s", e.service)
}

func (e attemptToJoinServiceEvent) Age() int {
	return e.age
}
