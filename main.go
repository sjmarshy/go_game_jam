package main

import (
	_ "fmt"
	tb "github.com/nsf/termbox-go"
	"os"
)

const helpMessage = "Honestly, there is no help yet. You're pretty much going to have to figure it out - I'm making this game up as I go along and the plan is to get somewhere within a weekend so who knows how well that will go. Probably not well at all, but maybe that's okay. it's turn based, dust is a good and a bad thing, and I only have a vague idea where I'm going with this."

func handleEvent(ev tb.Event, s state) state {

	var err error

	switch ev.Ch {
	case keyc:

		err, s = build(s, collector)

		if err != nil {
			s.message = err.Error()
		}
		break
	}

	return s
}

func turn(s state) state {

	s.message = ""

	ev := s.nextTurnEvent

	s = handleEvent(ev, s)

	s.dust = s.dust + uint(1+s.howManyBuilt(collectorName))
	s.nextTurnEvent = tb.Event{}

	return s
}

func loop(s state) {

	for {
		render(s)

		switch ev := tb.PollEvent(); ev.Type {
		case tb.EventKey:
			switch ev.Key {
			case tb.KeyEsc:
				os.Exit(0)
				break
			case tb.KeySpace:
				s = turn(s)
				break
			default:
				switch ev.Ch {
				case keyq:
					os.Exit(0)
					break
				case keyh:
					s.message = helpMessage
					render(s)
					break
				default:
					s.nextTurnEvent = ev
					break
				}
			}
		}
	}
}

func main() {

	err := tb.Init()

	if err != nil {
		panic(err)
	}

	defer tb.Close()

	s := state{}

	loop(s)
}
