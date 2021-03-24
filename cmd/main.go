package main

import (
	"encoding/json"
	"io/ioutil"
	blt "roguelike/bearlibterminal"
	"roguelike/data"
	"roguelike/game_core"
	"roguelike/loader"
	"runtime"
	"strconv"
	"strings"
)

func init() {
	blt.Open()
}

func getWindowSize() (int, int) {
	windowSizeS := blt.Get("window.size", "0x0")
	sizes := strings.Split(windowSizeS, "x")
	windowX, _ := strconv.Atoi(sizes[0])
	windowY, _ := strconv.Atoi(sizes[1])
	return windowX, windowY
}

func findPlayer(list []*game_core.Actor) *game_core.Actor {
	for _, v := range list {
		if v.Type == "PlayerBody" {
			return v
		}
	}
	return nil
}

func main() {
	runtime.LockOSThread()
	mapData, _ := ioutil.ReadFile("assets/map.json")
	mapObject := loader.Map{}
	_ = json.Unmarshal(mapData, &mapObject)
	objects := mapObject.ToActorList()
	player := findPlayer(objects)
	windowX, windowY := getWindowSize()
	camera := game_core.Camera{
		WindowSize: &data.Vector{
			X: windowX,
			Y: windowY,
		},
		Focus: player,
	}
	game := game_core.NewGame(&camera, player, objects)
	handleInput := game_core.CreateInputHandler(player)
	for {
		game.Render()
		blt.Refresh()
		key := blt.Read()
		if key == blt.TK_CLOSE {
			break
		}
		handleInput(key)
		game.UpdateSystems()
		blt.Clear()
	}

	blt.Close()
}
