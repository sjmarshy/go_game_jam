package main

import (
	"fmt"
	tb "github.com/nsf/termbox-go"
	"strings"
)

func getDescription(ev tb.Event) string {

	r := ""

	switch ev.Ch {
	case keyc:
		r = "build Collector"
	}

	return r
}

func drawString(x, y int, s string) {

	var sts []string

	if strings.ContainsRune(s, '\n') {

		sts = strings.Split(s, "\n")
	} else {

		sts = []string{s}
	}

	for i, st := range sts {

		for j, r := range st {

			nx := x + (j - 1)
			tb.SetCell(nx, y+i, r, tb.ColorBlack, tb.ColorWhite)
		}
	}
}

func addNewlines(s string, l int) string {

	ns := ""

	for i, c := range s {

		i = i + 1
		st := string(c)

		if i%l == 0 {
			ns = fmt.Sprint(ns, st, "\n")
		} else {
			ns = fmt.Sprint(ns, st)
		}
	}

	return ns
}

func drawRectangle(x, y, w, h int) {

	for i := x; i < w+x; i++ {
		for j := y; j < h+y; j++ {

			tb.SetCell(i, j, ' ', tb.ColorBlack, tb.ColorWhite)
		}
	}
}

func drawMessage(s string) {

	ns := addNewlines(s, 80)
	w, h := tb.Size()
	wc := w / 2
	hc := h / 2

	rectW := 82
	rectH := strings.Count(ns, "\n") + 3

	rectX := wc - rectW/2
	rectY := hc - rectH/2

	drawRectangle(rectX, rectY, rectW, rectH)
	drawString(rectX+2, rectY+1, ns)
}

func render(s state) {

	tb.Clear(tb.ColorDefault, tb.ColorDefault)

	w, h := tb.Size()

	drawString(1, 1, fmt.Sprint("dust: ", s.dust))

	if s.panelFound && s.highestDust >= 10 {
		drawString(1, 2, fmt.Sprint("[c]ollectors(", collector.currentCost(s), "): ", s.howManyBuilt("Collector")))
	}

	helpMessage := "hit [space] to take a turn, [h] for help, [q] or [Esc] to quit"
	nextTurnMessage := fmt.Sprint("next turn:", getDescription(s.nextTurnEvent))

	drawString((w/2)-(len(helpMessage)/2), h-1, helpMessage)
	drawString((w/2)-(len(nextTurnMessage)/2), h-2, nextTurnMessage)

	if len(s.message) > 0 {
		drawMessage(s.message)
	}

	err := tb.Flush()

	if err != nil {
		panic(err)
	}
}
