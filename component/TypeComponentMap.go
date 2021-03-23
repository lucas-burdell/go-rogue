package component

var TypeComponentMap = map[string]func() []BaseComponent{
	"Player": func() []BaseComponent {
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
