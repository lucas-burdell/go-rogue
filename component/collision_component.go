package component

var CollisionComponentName = "Collision"

type CollisionComponent struct {
	IsCollidable bool
}

func (c *CollisionComponent) Type() string {
	return CollisionComponentName
}
