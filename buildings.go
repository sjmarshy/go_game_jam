package main

import "fmt"

type cantAffordError struct {
	msg string
}

func (e *cantAffordError) Error() string {
	return e.msg
}

type effect struct {
	target    string
	action    string
	intensity int
}

type building struct {
	name     string
	baseCost uint
	effect   effect
}

const collectorName = "Collector"

var collector building = building{
	name:     collectorName,
	baseCost: 10,
	effect: effect{
		"dust",
		"inc",
		1,
	},
}

func build(s state, b building) (error, state) {

	if s.canAfford(b) {
		decreaseDust(&s, b.baseCost)
		addBuilding(&s, b)
	} else {

		return &cantAffordError{fmt.Sprint(b.name, " costs ", b.baseCost, " to build")}, s
	}

	return nil, s
}
