package main

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/awesome-gocui/gocui"
)

const NumGoroutines = 1

var (
	active = 0

	CrabWorkers = []*Crab{{name: "peacrab", cost: 5, production: 1, count: 1}, {name: "Sand Crab", cost: 10, production: 5}, {name: "King Crab", cost: 100, production: 50}, {name: "King Crab3", cost: 100, production: 100}, {name: "King Crab4", cost: 100, production: 100}, {name: "King Crab5", cost: 100, production: 100}}
	done        = make(chan struct{})
	wg          sync.WaitGroup

	mu         sync.Mutex // protects ctr
	bankAmount int
	income     = 1
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
	bankAmount = 0

	go counter(g, &income, &bankAmount)

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
}

func increaseIncome(amount int) {

	income += amount

}
func SpendMoney(cost int) bool {
	currentBank := &bankAmount
	if cost < *currentBank {
		bankAmount -= cost
		return true
	}
	return false
}

func counter(g *gocui.Gui, income *int, bankAmount *int) {
	defer wg.Done()

	for {
		select {
		case <-done:
			return
		case <-time.After(500 * time.Millisecond):
			mu.Lock()
			*bankAmount += *income
			mu.Unlock()

			showWorkers(g)

			showCurrency(g)
		}
	}
}
