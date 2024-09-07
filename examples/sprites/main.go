package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const UGENE = "../../assets/ugene.png"
const DOOR = "../../assets/door.png"

func main() {

	game := boom.NewGame(800, 600, "Sprites")
	game.SetBgColor(rl.Black)
	scene := game.GetCurrentScene()

	rect := boom.Rectangle(300, 100, 50, 50, rl.Red, rl.Blue, 2)
	circle := boom.Circle(100, 100, 50, rl.Red, rl.Blue, 2)
	sprite := boom.Sprite(100, 100, UGENE)

	scene.Add(sprite, circle, rect)

	game.Run()

}
