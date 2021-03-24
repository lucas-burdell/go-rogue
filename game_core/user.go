package game_core

import (
	blt "roguelike/bearlibterminal"
	"roguelike/component"
	"roguelike/data"
)

func CreateInputHandler(player *Actor) func(key int) {
	movement := player.FindComponentByName("Movement").(*component.MovementComponent)
	return func(key int) {
		switch key {
		case blt.TK_D:
			movement.Velocity = &data.Vector{
				X: 1,
			}
		case blt.TK_A:
			movement.Velocity = &data.Vector{
				X: -1,
			}
		case blt.TK_W:
			movement.Velocity = &data.Vector{
				Y: -1,
			}
		case blt.TK_S:
			movement.Velocity = &data.Vector{
				Y: 1,
			}
		}
	}
}
