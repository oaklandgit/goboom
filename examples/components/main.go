package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	game := boom.NewGame(800, 600, "Components")
	game.SetBgColor(rl.Red)
	scene := game.GetCurrentScene()

	player := boom.NewRectangle(0, 0, 50, 50, rl.Blank)
	scene.Add(player)

	game.Run()
	
}




