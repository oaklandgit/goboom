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

	// Game objects can be created directly
	rect := boom.Rectangle(300, 100, 50, 50, rl.Pink, rl.Blue, 8)
	circle := boom.Circle(10, 20, 30, rl.Red, rl.Yellow, 2)
	sprite := boom.Sprite(100, 100, UGENE)
	poly := boom.Polygon(120, 220, "12 13 1 8.5 22 32 101 0", true, rl.Green, rl.Yellow, 2)
	regpoly := boom.RegPolygon(200, 200, 60, 6, rl.Blue, rl.Yellow, 2)

	// Add new components to any existing game object
	multi := boom.Polygon(120, 220, "12 13 1 8.5 22 32 101 0", true, rl.Red, rl.Yellow, 2)
	multi.AddComponent(boom.NewSpriteComp(DOOR))


	// or create a new game object with multiple components
	test := boom.Create(
		200, 300,
		boom.NewRectComp(30, 40, rl.Green, rl.Yellow, 2),
		boom.NewSpriteComp(UGENE),
		boom.NewPolyComp("12 13 1 8.5 22 32 101 10", true, rl.Blue, rl.Purple, 2),
		boom.NewVelocityComp(0.5, 0),
	)

	anim := boom.Sprite(400, 400, UGENE, DOOR)

	anim.Get("sprite").(*boom.SpriteComp).AddAnim("idle", []int{0, 1}, 2, true)
	anim.Get("sprite").(*boom.SpriteComp).Play("idle")

	test.SetAngle(12)
	rect.SetAngle(30)
	circle.SetScale(2, 1)
	sprite.SetScale(2, 2)
	regpoly.SetScale(2, 3)

	test.AddComponent(boom.NewInputComp())

	turnLeft := boom.NewInput(rl.KeyLeft, boom.KeyDown, func() {
		test.AddAngle(-1)
	})

	turnRight := boom.NewInput(rl.KeyRight, boom.KeyDown, func() {
		test.AddAngle(1)
	})

	test.AddComponent(boom.NewInputComp(turnLeft, turnRight))

	circle.With(boom.NewVelocityComp(0, 0))
	circle.Get("velocity").(*boom.VelocityComp).SetVelocity(2, 1)


	scene.Add(circle, rect, sprite, poly, regpoly, multi, test, anim)

	game.Run()

}
