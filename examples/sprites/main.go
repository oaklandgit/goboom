package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const COMPUTER = "../../assets/computer.png"
const DOOR = "../../assets/door.png"

func main() {

	game := boom.NewGame(800, 600, "Sprites")
	game.SetBgColor(rl.Black)


	scene := game.GetCurrentScene()

	// circle := boom.NewCircleObj(100, 100, 50, rl.Red, rl.Blue, 2)
	// scene.Add(circle)

	circle := boom.NewGameObject()
	c := boom.NewCircleComp(50)
	c.FillColor = rl.Red

	circle.AddComponent(c)

	scene.Add(circle)

	player := boom.NewGroup(0, 0)
	player.SetOrigin(0.5, 0.5)
	player.SetScale(2, 2)
	player.SetAngle(45)
	player.SetXY(scene.CenterX(), scene.CenterY())
	player.AddComponent(boom.NewSpriteComp(COMPUTER, DOOR))

	scene.Add(player)

	game.Run()

}
