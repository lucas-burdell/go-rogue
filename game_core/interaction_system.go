package game_core

type interactionSystem struct {
	game *Game
}

func (pis *interactionSystem) Type() string {
	return "Interaction"
}

func NewInteractionSystem(game *Game) *interactionSystem {
	return &interactionSystem{
		game: game,
	}
}

func (i *interactionSystem) Evaluate(actor *Actor) {

}
