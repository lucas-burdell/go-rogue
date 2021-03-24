package game_core

import (
	"fmt"
	blt "roguelike/bearlibterminal"
	"roguelike/data"
)

type Camera struct {
	Focus      *Actor
	WindowSize *data.Vector
}

func (c *Camera) renderChildObject(a *Actor, parent *Actor) {
	x := (c.Focus.Position.X - (parent.Position.X - a.Position.X)) * 4
	y := (c.Focus.Position.Y - (parent.Position.Y - a.Position.Y)) * 4
	symbol := fmt.Sprintf("[font=%s]%s[/font]", a.TileMap, string(a.Symbol))
	blt.Layer(a.Layer)
	blt.Print(((c.WindowSize.X / 2) - x), ((c.WindowSize.Y / 2) - y), symbol)
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
	symbol := fmt.Sprintf("[font=%s]%s[/font]", a.TileMap, string(a.Symbol))

	x := (c.Focus.Position.X - a.Position.X) * 4
	y := (c.Focus.Position.Y - a.Position.Y) * 4
	blt.Layer(a.Layer)
	blt.Print(((c.WindowSize.X / 2) - x), ((c.WindowSize.Y / 2) - y), symbol)
	if a.Attachments != nil {
		for _, v := range a.Attachments {
			c.renderChildObject(v, a)
		}
	}
}
