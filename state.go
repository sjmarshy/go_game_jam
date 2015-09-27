package main

import tb "github.com/nsf/termbox-go"

type state struct {
	nextTurnEvent tb.Event
	dust          float32
	disks         float32
	highestDust   float32
	message       string
	panelFound    bool
	firsts        map[string]bool
	buildings     []building
}

func (s state) canAfford(b building) bool {

	cost := b.currentCost(s)

	return s.dust >= cost
}

func decreaseDust(s *state, by float32) {
	s.dust -= by
}

func addBuilding(s *state, b building) {
	s.buildings = append(s.buildings, b)
}

func (s state) howManyBuilt(name string) int {

	c := 0

	for _, b := range s.buildings {
		if b.name == name {
			c = c + 1
		}
	}

	return c
}
