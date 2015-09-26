package main

import tb "github.com/nsf/termbox-go"

type state struct {
	nextTurnEvent tb.Event
	dust          int
	collectors    int
	message       string
}
