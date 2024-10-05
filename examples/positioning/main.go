package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	game := boom.NewGame(800, 600, "Positioning")
	game.SetBgColor(rl.Black)
	scene := game.GetCurrentScene()


	// Create a rectangle
	// shape := boom.Rectangle(0, 0, 100, 100, rl.Red, rl.White, 2)
	// shape := boom.Circle(0, 0, 50, rl.Blue, rl.White, 2)
	// shape := boom.RegPolygon(0, 0, 50, 6, rl.Green, rl.White, 2)
	shape := boom.Polygon(0, 0, "12 13 1 8.5 22 32 101 0", true, rl.Yellow, rl.White, 2)

	scene.Add(shape)

	game.Run()

}