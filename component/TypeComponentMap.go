package component

var TypeComponentMap = map[string]func() []BaseComponent{
	"PlayerBody": func() []BaseComponent {
		return []BaseComponent{
			&MovementComponent{},
			&CollisionComponent{
				IsCollidable: true,
			},
		}
	},
	"Wall": func() []BaseComponent {
		return []BaseComponent{
			&CollisionComponent{
				IsCollidable: true,
			},
		}
	},
}
