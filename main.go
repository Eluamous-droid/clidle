package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/awesome-gocui/gocui"
)

const NumGoroutines = 1
const ControlView = "ControlView"
const CurrencyView = "CurrencyView"
const UpgradeView = "UpgradeView"

var (
	viewArr = []string{ControlView,CurrencyView, UpgradeView, "v4"}
	active  = 0

	done = make(chan struct{})
	wg   sync.WaitGroup

	mu  sync.Mutex // protects ctr
	ctr = 0
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	initKeybinds(g)

	go counter(g)

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
}


func counter(g *gocui.Gui) {
	defer wg.Done()

	for {
		select {
		case <-done:
			return
		case <-time.After(500 * time.Millisecond):
			mu.Lock()
			n := ctr
			ctr++
			mu.Unlock()

			g.Update(func(g *gocui.Gui) error {
				v, err := g.View(CurrencyView)
				if err != nil {
					return err
				}
				v.Clear()
				fmt.Fprintln(v, n)
				return nil
			})
		}
	}
}
