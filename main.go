package main

import (
	tb "github.com/nsf/termbox-go"
	"os"
)

func handleEvent(ev tb.Event, s state) state {

	switch ev.Ch {
	case keyc:
		s.collectors++
		break
	}

	return s
}

func turn(s state) state {

	ev := s.nextTurnEvent

	s = handleEvent(ev, s)

	incDustBy := 1 + s.collectors
	s.dust = s.dust + incDustBy
	s.nextTurnEvent = tb.Event{}

	// messages only last for a single turn
	s.message = ""

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
					s.message = "super long help message just to see if the line splits like it should. I mean it wont yet because I havent done the draw function right but you know, you know buddy. You know."
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
