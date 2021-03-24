package game_core

import (
	"roguelike/component"
	"roguelike/data"
)

type movementSystem struct {
	game *Game
}

func (pis *movementSystem) Type() string {
	return "Movement"
}

func NewMovementSystem(game *Game) *movementSystem {
	return &movementSystem{
		game: game,
	}
}

func (m *movementSystem) Evaluate(actor *Actor) {
	if actor == nil {
		return
	}
	movement := actor.FindComponentByName("Movement")
	if movement == nil {
		return
	}
	movementComp := movement.(*component.MovementComponent)
	if movementComp.Velocity == nil {
		return
	}
	intendedPosition := data.Vector{
		X: actor.Position.X + movementComp.Velocity.X,
		Y: actor.Position.Y + movementComp.Velocity.Y,
	}
	collidables, collisionMap := m.game.GetAllActorsByComponentName("Collision")
	for _, v := range collidables {
		if collisionMap[v].(*component.CollisionComponent).IsCollidable && *v.Position == intendedPosition {
			movementComp.Velocity = &data.Vector{}
			return
		}
	}
	actor.Position = &intendedPosition
	movementComp.Velocity = &data.Vector{}
}
