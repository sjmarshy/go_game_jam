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

func handleDiskEffect(s state, e effect) state {
	switch e.action {
	case "inc":
		s.disks += e.intensity
	}

	return s
}

func handleBuildingEffect(s state, b building) state {

	ef := b.effect

	switch ef.target {
	case "dust":
		s = handleDustEffect(s, ef)
		break
	case "disks":
		s = handleDiskEffect(s, ef)
		break
	}

	return s
}

var collector building = building{
	name:     "Collector",
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

var searcher building = building{
	name:     "Searcher",
	baseCost: 50,
	costAdjust: func(amt int, baseCost float32) float32 {
		return float32(amt)*1.8 + baseCost
	},
	effect: effect{
		"disks",
		"inc",
		0.1,
	},
}
