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
	if v, err := g.SetView(TitelView, 0, 0, maxX/2, bannerHeight, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
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

	if v, err := g.SetView(CurrencyView, maxX/2-1, 0, maxX-1, bannerHeight, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Frame = false;
		v.Wrap = true
		showCurrency(g)
	}
	if _, err := g.SetView(UpgradeView, 0, maxY/2-1, maxX/2-1, maxY-1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		showShop(g)
	}
	if _, err := g.SetView(WorkerView, maxX/2, bannerHeight+1, maxX-1, maxY-1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		showWorkers(g)
	}
	return nil
}


func showCurrency(g *gocui.Gui){

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(CurrencyView)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, "Plankton count: " + strconv.Itoa(bankAmount))
		return nil
	})

}


func showWorkers(g *gocui.Gui) {

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(WorkerView)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, "Press number to purchase")
		for i:=0; i<3; i++{
			c := CrabWorkers[i]
			fmt.Fprintln(v, strconv.Itoa(i + 1) + ": " + c.name + ": Cost: " + strconv.Itoa(c.cost) + ", Production: " + strconv.Itoa(c.production) + ", Count: " + strconv.Itoa(c.count))
			
			
		}
		return nil
	})

}


func showShop(g *gocui.Gui){

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(WorkerView)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, "Press letter to purchase")
		return nil
	})
}


func showInstructions(g *gocui.Gui){

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(WorkerView)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, "Put instructions here")
		fmt.Fprintln(v, "Figure out controls first")
		return nil
	})


}
