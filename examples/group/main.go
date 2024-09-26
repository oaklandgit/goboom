package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	game := boom.NewGame(800, 600, "Group Example")
	scene := game.GetCurrentScene()

	// A function that makes an object with a rectangle and an circle
	makeThingy := func() *boom.GameObject {
		thingy := boom.NewGroup(
			0, 0,
			boom.Rectangle(0, 0, 100, 100, rl.Blue, rl.Black, 2),
			boom.Ellipse(0, 0, 100, 100, rl.Red, rl.Black, 2),
		)

		return thingy
	}

	// make a 3x5 grid with a 4px gap between each object
	groupOfThingies := boom.GridArray(3, 5, 4, makeThingy)

	// boom.PutCenter(scene, groupOfThingies, 0, 0)

	scene.Add(groupOfThingies)

	game.Run()

}