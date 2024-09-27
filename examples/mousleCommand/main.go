package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)


func main() {

	game := boom.NewGame(800, 600, "Mousle Command")
	game.SetBgColor(rl.Black)
	scene := game.GetCurrentScene()

	// Create the crosshairs
	crosshairs := boom.NewGroup(0, 0)
	crosshairs.Add(boom.Rectangle(0, 0, 20, 2, rl.White, rl.White, 1))
	crosshairs.Add(boom.Rectangle(0, 0, 2, 20, rl.White, rl.White, 1))

	// Mouse component
	mouse := boom.NewMouseComp()
	mouse.OnMove(func(x, y float32) {
		crosshairs.SetXY(x, y)
	})
	crosshairs.AddComponent(mouse)

	scene.Add(crosshairs)

	boom.PutCenter(scene, crosshairs, 0, 0)

	game.Run()

}



