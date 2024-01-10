package main

import (
	"github.com/awesome-gocui/gocui"
)

func initKeybinds(g *gocui.Gui) error{

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return nil
	}
	if err := g.SetKeybinding("stdin", '1', gocui.ModNone, buyWorker()); err != nil {
		return nil
	}
	return nil

}

func buyWorker(c crab){
	c.buyCrab()
}



func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
