package main

import (
	"github.com/awesome-gocui/gocui"
)

var hexValues = [10]rune{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x30}

func initKeybinds(g *gocui.Gui) error{

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return nil
	}
	if err := g.SetKeybinding("", hexValues[0], gocui.ModNone,func(g *gocui.Gui, v *gocui.View) error {
		buyWorker(CrabWorkers[0], g)
		return nil
	} ); err != nil {
		return err
	}
	if err := g.SetKeybinding("", hexValues[1], gocui.ModNone,func(g *gocui.Gui, v *gocui.View) error {
		buyWorker(CrabWorkers[1], g)
		return nil
	} ); err != nil {
		return err
	}
	if err := g.SetKeybinding("", hexValues[2], gocui.ModNone,func(g *gocui.Gui, v *gocui.View) error {
		buyWorker(CrabWorkers[2], g)
		return nil
	} ); err != nil {
		return err
	}
	
	return nil

}

func buyWorker(c *Crab, g *gocui.Gui) error{
	if SpendMoney(c.cost){
		c.buyCrab()
	}
	return nil

}



func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
