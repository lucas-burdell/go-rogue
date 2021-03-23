package game_core

type BaseSystem interface {
	Type() string
	Evaluate(actor *Actor)
}
