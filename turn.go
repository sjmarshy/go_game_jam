package main

import tb "github.com/nsf/termbox-go"

func handleEvent(ev tb.Event, s state) state {

	var err error

	switch ev.Ch {
	case keyc:
		if s.panelFound {
			err, s = build(s, collector)

			if err != nil {
				s.message = err.Error()
			}
		}
		break
	}

	return s
}

func turn(s state) state {

	s.message = ""

	ev := s.nextTurnEvent

	s = handleEvent(ev, s)

	for _, b := range s.buildings {
		s = handleBuildingEffect(s, b)
	}

	s.dust += 1
	s.nextTurnEvent = tb.Event{}

	if s.dust > s.highestDust {
		s.highestDust = s.dust
	}

	if s.highestDust == 8 {
		s.panelFound = true
		s.message = "You're not sure why you're scrabbling around in the dust, but it seems to be something to do. While you're getting your hands into some particularly grainy feeling dust, you spot a blue glow coming from under a nearby low dune. Crawling over for a closer look, you see it's a rectangle, with a patch of light under some...glass? It shows the number 8. You keep on scrabbling in the dust"
	}

	if s.highestDust == 10 {
		s.message = "something flickers on the screen as your're digging, and you see a set of characters appear. Something about 'collectors'. You figure it might be worth a prod"
	}

	if s.howManyBuilt(collectorName) == 1 && s.firstCollector {
		s.firstCollector = false
		s.message = "the dust you seem to have collected under your fingernails clumps together in front of your eyes and, with a faint blue sheen, lands in front of you. The dust around it seems to be attracted to it"
	}

	return s
}
