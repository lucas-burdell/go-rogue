package game_core

import "rougelike/component"

type Game struct {
	Camera  *Camera
	Player  *Actor
	Actors  []*Actor
	systems []BaseSystem
}

func NewGame(cam *Camera, player *Actor, actors []*Actor) *Game {
	output := &Game{
		Camera:  cam,
		Actors:  actors,
		Player:  player,
		systems: []BaseSystem{},
	}
	movement := NewMovementSystem(output)
	output.addSystem(movement)

	return output
}

func (g *Game) addSystem(newSystem BaseSystem) {
	g.systems = append(g.systems, newSystem)
}

func (g *Game) GetAllActorsByComponentName(componentName string) ([]*Actor, map[*Actor]component.BaseComponent) {
	output := make([]*Actor, 0)
	outputMap := map[*Actor]component.BaseComponent{}
	for _, v := range g.Actors {
		comp := v.FindComponentByName(componentName)
		if comp != nil {
			output = append(output, v)
			outputMap[v] = comp
		}
	}
	return output, outputMap
}

func (g *Game) Render() {
	for _, v := range g.Actors {
		g.Camera.RenderObject(v)
	}
}

func (g *Game) UpdateSystems() {
	for _, system := range g.systems {
		for _, v := range g.Actors {
			system.Evaluate(v)
		}
	}
}
