package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func objectGenerator() func() *boom.GameObject {
	objects := []*boom.GameObject{
		boom.Rectangle(0, 0, 100, 100, rl.Red, rl.White, 2),
		boom.Circle(0, 0, 50, rl.Blue, rl.White, 2),
		boom.RegPolygon(0, 0, 50, 6, rl.Green, rl.White, 2),
		boom.Polygon(0, 0, "12 13 1 8.5 22 32 101 0", true, rl.Yellow, rl.White, 2),
	}
	index := 0

	return func() *boom.GameObject {
		if index >= len(objects) {
			return nil // Return nil if no more objects are available
		}
		obj := objects[index]
		index++
		return obj
	}
}

func main() {

	game := boom.NewGame(800, 600, "Positioning")
	game.SetBgColor(rl.Black)
	scene := game.GetCurrentScene()

	group := boom.GridArray(2, 2, 100, objectGenerator())
	

	// group := boom.NewGroup(0, 0, rect, circ, rpoly, poly)

	// collide := boom.NewCollideComp(boom.CollisionRect{})
	// shape.AddComponents(collide)

	scene.Add(group)

	game.Run()

}