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
	intensity float32
}

type building struct {
	name       string
	baseCost   float32
	costAdjust func(amt int, baseCost float32) float32
	effect     effect
}

func (b building) currentCost(s state) float32 {
	return b.costAdjust(s.howManyBuilt(b.name), b.baseCost)
}

const collectorName = "Collector"

var collector building = building{
	name:     collectorName,
	baseCost: 10,
	costAdjust: func(amt int, baseCost float32) float32 {

		return float32(amt)*1.5 + baseCost
	},
	effect: effect{
		"dust",
		"inc",
		1,
	},
}

func build(s state, b building) (error, state) {

	if s.canAfford(b) {
		decreaseDust(&s, b.currentCost(s))
		addBuilding(&s, b)
	} else {

		return &cantAffordError{fmt.Sprint(b.name, " costs ", b.currentCost(s), " to build")}, s
	}

	return nil, s
}

func handleDustEffect(s state, e effect) state {

	switch e.action {
	case "inc":
		s.dust += e.intensity
	}

	return s
}

func handleBuildingEffect(s state, b building) state {

	ef := b.effect

	switch ef.target {
	case "dust":
		s = handleDustEffect(s, ef)
		break
	}

	return s
}
