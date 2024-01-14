package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/awesome-gocui/gocui"
)

const ControlView = "ControlView"
const CurrencyView = "CurrencyView"
const UpgradeView = "UpgradeView"
const TitelView = "TitelView"
const WorkerView = "Workerview"

var viewArr = []string{ControlView,CurrencyView, UpgradeView, "v4"}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	bannerHeight := 4
	if v, err := g.SetView(TitelView, 0, 0, maxX, bannerHeight, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Wrap = true
		v.Frame = false
		fmt.Fprintln(v, " __   __        __                __   __   __      ")
		fmt.Fprintln(v, "/  ` |__)  /\\  |__)    |     /\\  / _` /  \\ /  \\ |\\ |")
		fmt.Fprintln(v, "\\__, |  \\ /~~\\ |__)    |___ /~~\\ \\__> \\__/ \\__/ | \\|" )

	}
	if v, err := g.SetView(ControlView, 0, bannerHeight+1, maxX/2-1, maxY/2-1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Wrap = true

	}

	if v, err := g.SetView(CurrencyView, maxX/2-1, bannerHeight+1, maxX-1, maxY/2-1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, bankAmount)
	}
	if v, err := g.SetView(UpgradeView, 0, maxY/2-1, maxX/2-1, maxY-1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		fmt.Fprint(v, "Press TAB to change current view")
	}
	if _, err := g.SetView(WorkerView, maxX/2, maxY/2, maxX-1, maxY-1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		showWorkers(g)
	}
	return nil
}


func showWorkers(g *gocui.Gui) {

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(WorkerView)
		if err != nil {
			return err
		}
		v.Clear()
		for i:=0; i<3; i++{
			c := CrabWorkers[i]
			fmt.Fprintln(v, c.name + ": Cost: " + strconv.Itoa(c.cost) + ", Production: " + strconv.Itoa(c.production) + ", Count: " + strconv.Itoa(c.count))
			
			
		}
		return nil
	})

}
