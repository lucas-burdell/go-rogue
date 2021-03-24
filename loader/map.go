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
	Objects []*object `json:"objects"`
}

type object struct {
	Tile       int         `json:"gid"`
	Type       string      `json:"type"`
	X          float64     `json:"x"`
	Y          float64     `json:"y"`
	Visible    bool        `json:"visible"`
	ID         int         `json:"id"`
	Properties *[]property `json:"properties"`
}

type property struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type objectActorPair struct {
	Object *object
	Actor  *game_core.Actor
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
			TileMap:    "tiles",
			Type:       tileType,
			Components: loadComponents(tileType),
		})
	}
	return output
}

func (l *layer) toObjectActors(layerid int) map[int]objectActorPair {
	objectMap := map[int]objectActorPair{}
	for _, v := range l.Objects {
		x := int(v.X / 16)
		y := int(v.Y / 16)
		actor := &game_core.Actor{
			Position: &data.Vector{
				X: x,
				Y: y,
			},
			Symbol:      rune((v.Tile - 334) + 0x1000),
			Layer:       layerid,
			TileMap:     "chars",
			Type:        v.Type,
			Components:  loadComponents(v.Type),
			Attachments: []*game_core.Actor{},
		}
		objectMap[v.ID] = objectActorPair{
			Object: v,
			Actor:  actor,
		}
	}
	return objectMap
}

func resolveParents(objectMap map[int]objectActorPair) []*game_core.Actor {
	output := make([]*game_core.Actor, 0)
	for _, pair := range objectMap {
		if pair.Object.Properties == nil {
			output = append(output, pair.Actor)
			continue
		}
		parent := -1
		for _, v := range *pair.Object.Properties {
			if v.Name == "Parent" {
				parent = int(v.Value.(float64))
			}
		}
		if parent == -1 {
			output = append(output, pair.Actor)
			continue
		}
		pair.Actor.Position = &data.Vector{
			X: objectMap[parent].Actor.Position.X - pair.Actor.Position.X,
			Y: objectMap[parent].Actor.Position.Y - pair.Actor.Position.Y,
		}
		objectMap[parent].Actor.Attachments = append(objectMap[parent].Actor.Attachments, pair.Actor)
		output = append(output, pair.Actor)
	}
	return output
}

func (tileMap *Map) ToActorList() []*game_core.Actor {
	backgroundLayer := findLayerByName(tileMap.Layers, "Background")
	wallsLayer1 := findLayerByName(tileMap.Layers, "Walls_1")
	wallsLayer2 := findLayerByName(tileMap.Layers, "Walls_2")
	objectsLayer := findLayerByName(tileMap.Layers, "Objects")

	background := backgroundLayer.toTileActors(0, "Background")
	walls := wallsLayer1.toTileActors(1, "Wall")
	walls = append(walls, wallsLayer2.toTileActors(2, "Wall")...)
	objectMap := objectsLayer.toObjectActors(3)
	objects := resolveParents(objectMap)

	output := append(background, walls...)
	output = append(output, objects...)
	return output
}
