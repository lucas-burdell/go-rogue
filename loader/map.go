package loader

import (
	"roguelike/component"
	"roguelike/data"
	"roguelike/game_core"
)

type Map struct {
	Height int      `json:"height"`
	Width  int      `json:"width"`
	Layers []*layer `json:"layers"`
}

type layer struct {
	Height  int       `json:"height"`
	Width   int       `json:"width"`
	Name    string    `json:"name"`
	LayerID int       `json:"id"`
	Tiles   []int     `json:"data"`
	Objects *[]object `json:"objects"`
}

type object struct {
	Tile    int     `json:"gid"`
	Type    string  `json:"type"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Visible bool    `json:"visible"`
}

func loadComponents(entityType string) []component.BaseComponent {
	if component.TypeComponentMap[entityType] == nil {
		return []component.BaseComponent{}
	}
	return component.TypeComponentMap[entityType]()
}

func findLayerByName(layers []*layer, name string) *layer {
	for _, v := range layers {
		if v.Name == name {
			return v
		}
	}
	return nil
}

func (l *layer) toTileActors(layerid int, tileType string) []*game_core.Actor {
	output := make([]*game_core.Actor, 0)
	for index, tile := range l.Tiles {
		x := index % (l.Width)
		y := index / (l.Width)
		if tile == 0 {
			continue
		}
		output = append(output, &game_core.Actor{
			Position: &data.Vector{
				X: x,
				Y: y,
			},
			Symbol:     rune(tile + 0x1000 - 1),
			Layer:      layerid,
			IsTile:     true,
			Type:       tileType,
			Components: loadComponents(tileType),
		})
	}
	return output
}

func (l *layer) toObjectActors(layerid int) []*game_core.Actor {
	output := make([]*game_core.Actor, len(*l.Objects))
	for index, v := range *l.Objects {
		x := 2
		y := 2
		output[index] = &game_core.Actor{
			Position: &data.Vector{
				X: x,
				Y: y,
			},
			Symbol:     rune(v.Tile + 0x1000 - 1),
			Layer:      layerid,
			IsTile:     true,
			Type:       v.Type,
			Components: loadComponents(v.Type),
		}
	}
	return output
}

func (tileMap *Map) ToActorList() []*game_core.Actor {
	backgroundLayer := findLayerByName(tileMap.Layers, "Background")
	wallsLayer := findLayerByName(tileMap.Layers, "Walls")
	objectsLayer := findLayerByName(tileMap.Layers, "Objects")

	background := backgroundLayer.toTileActors(0, "Background")
	walls := wallsLayer.toTileActors(1, "Wall")
	objects := objectsLayer.toObjectActors(2)

	output := append(background, walls...)
	output = append(output, objects...)
	return output
}
