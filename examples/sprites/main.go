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

	rect := boom.Rectangle(300, 100, 50, 50, rl.Pink, rl.Blue, 8)
	circle := boom.Circle(10, 20, 30, rl.Red, rl.Yellow, 2)
	sprite := boom.Sprite(100, 100, UGENE)

	rect.SetAngle(30)
	circle.SetScale(2, 1)
	sprite.SetScale(2, 2)

	scene.Add(circle, rect, sprite)

	game.Run()

}
