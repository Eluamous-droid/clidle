package main

import (
	"github.com/awesome-gocui/gocui"
)

func initKeybinds(g *gocui.Gui) error{

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return nil
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, changeView); err != nil {
		return err
	}
	return nil

}

func changeView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	active = nextIndex
	return nil
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
