package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	game := boom.NewGame(800, 600, "Empty Game")

	scene := game.GetCurrentScene()

	message := "HELLO WORLD!"

	poly := boom.NewPolygon(0, 0, rl.Red, true, "0 0 10 0 10 10 0 10")
	circle := boom.NewCircle(0, 0, 10, rl.Blue)
	rect := boom.NewRectangle(0, 0, 40, 40, rl.Orange)
	regpoly := boom.NewRegPoly(0, 0, 5, 50, rl.Pink)
	text := boom.NewText(0, 0, &message, rl.Black)
	scene.Add(circle, poly, rect, regpoly, text)

	game.Run()
}