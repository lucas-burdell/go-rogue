package component

import "rougelike/data"

var MovementComponentName = "Movement"

type MovementComponent struct {
	Velocity *data.Vector
}

func (c *MovementComponent) Type() string {
	return MovementComponentName
}
