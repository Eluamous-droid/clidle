package main

import (
	"errors"
	"fmt"

	"github.com/awesome-gocui/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(ControlView, 0, 0, maxX/2-1, maxY/2-1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Title = "v1 (editable)"
		v.Editable = true
		v.Wrap = true

		if _, err = setCurrentViewOnTop(g, ControlView); err != nil {
			return err
		}
	}

	if v, err := g.SetView(CurrencyView, maxX/2-1, 0, maxX-1, maxY/2-1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Title = "v2"
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, ctr)
	}
	if v, err := g.SetView(UpgradeView, 0, maxY/2-1, maxX/2-1, maxY-1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Title = "v3"
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "Press TAB to change current view")
	}
	if v, err := g.SetView("v4", maxX/2, maxY/2, maxX-1, maxY-1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
		v.Title = "v4 (editable)"
		v.Editable = true
	}
	return nil
}
