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

type joinedServiceEvent struct {
	service career
	age     int
}

func (e joinedServiceEvent) Name() string {
	return fmt.Sprintf("began career: %s", e.service)
}

func (e joinedServiceEvent) Age() int {
	return e.age
}

type failedToJoinServiceEvent struct {
	service career
	age     int
}

func (e failedToJoinServiceEvent) Name() string {
	return fmt.Sprintf("rejected from: %s", e.service)
}

func (e failedToJoinServiceEvent) Age() int {
	return e.age
}

type draftedEvent struct {
	service career
	age     int
}

func (e draftedEvent) Name() string {
	return fmt.Sprintf("drafted into: %s", e.service)
}

func (e draftedEvent) Age() int {
	return e.age
}
