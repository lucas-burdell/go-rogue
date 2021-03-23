package game_core

import (
	blt "bearlibterminal"
	"fmt"
	"rougelike/data"
)

type Camera struct {
	Focus      *Actor
	WindowSize *data.Vector
}

func (c *Camera) RenderObject(a *Actor) {
	if a == nil {
		panic("wtf did you do")
	}
	if a.Position.X > c.Focus.Position.X+(c.WindowSize.X/2) || a.Position.X < c.Focus.Position.X-(c.WindowSize.X/2) {
		return
	}
	if a.Position.Y > c.Focus.Position.Y+(c.WindowSize.Y/2) || a.Position.Y < c.Focus.Position.Y-(c.WindowSize.Y/2) {
		return
	}
	symbol := string(a.Symbol)
	if a.IsTile {
		symbol = fmt.Sprintf("[font=tiles]%s[/font]", symbol)
	}
	blt.Layer(a.Layer)
	x := ((c.WindowSize.X / 2) - (c.Focus.Position.X - a.Position.X)) * 64
	y := ((c.WindowSize.Y / 2) - (c.Focus.Position.Y - a.Position.Y)) * 64
	blt.Print(x, y, symbol)
}
