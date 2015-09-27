package main

import (
	_ "fmt"
	tb "github.com/nsf/termbox-go"
	"os"
)

const helpMessage = "Honestly, there is no help yet. You're pretty much going to have to figure it out - I'm making this game up as I go along and the plan is to get somewhere within a weekend so who knows how well that will go. Probably not well at all, but maybe that's okay. it's turn based, dust is a good and a bad thing, and I only have a vague idea where I'm going with this."

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

	s := state{
		firsts: map[string]bool{
			collector.name: true,
			searcher.name:  true,
		},
	}

	loop(s)
}
