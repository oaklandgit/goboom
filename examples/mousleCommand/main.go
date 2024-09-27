package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)


func main() {

	game := boom.NewGame(800, 600, "Mousle Command")
	game.SetBgColor(rl.Black)
	scene := game.GetCurrentScene()

	// A blank game object to hold the crosshairs
	crosshairs := boom.NewGameObject()

	// Mouse component
	mouse := boom.NewMouseComp()
	mouse.SetCursor(rl.MouseCursorCrosshair)
	mouse.OnMove(func(x, y float32) {
		crosshairs.SetXY(x, y)
	})

	// splosions
	mouse.OnClick(func() {
		boom.NewExplosion(scene, crosshairs.GetX(), crosshairs.GetY(), 3, 12, 1, 3, func() *boom.GameObject {
			return boom.Rectangle(0, 0, 7, 5, rl.Blank, rl.White, 1)
		})
	})

	crosshairs.AddComponent(mouse)

	scene.Add(crosshairs)

	boom.PutCenter(scene, crosshairs, 0, 0)

	game.Run()

}



