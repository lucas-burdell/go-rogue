package game_core

import (
	"roguelike/component"
	"roguelike/data"
)

type Actor struct {
	Position   *data.Vector
	Symbol     rune
	Layer      int
	IsTile     bool
	Type       string
	Components []component.BaseComponent
}

func (a *Actor) FindComponentByName(name string) component.BaseComponent {
	for _, v := range a.Components {
		if v.Type() == name {
			return v
		}
	}
	return nil
}
