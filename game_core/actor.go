package game_core

import (
	"roguelike/component"
	"roguelike/data"
)

type Actor struct {
	Position    *data.Vector
	Symbol      rune
	Layer       int
	TileMap     string
	Type        string
	Components  []component.BaseComponent
	Attachments []*Actor
}

func (a *Actor) FindComponentByName(name string) component.BaseComponent {
	for _, v := range a.Components {
		if v.Type() == name {
			return v
		}
	}
	return nil
}
