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
	ellipse := boom.Ellipse(400, 200, 50, 10, rl.Red, rl.Yellow, 2)
	sprite := boom.Sprite(80, 120, UGENE)
	sprite.SetSize(100, 200)
	poly := boom.Polygon(120, 220, "12 13 1 8.5 22 32 101 0", true, rl.Green, rl.Yellow, 2)
	regpoly := boom.RegPolygon(200, 200, 60, 6, rl.Blue, rl.Yellow, 2)

	// Add new components to any existing game object
	multi := boom.Polygon(120, 220, "12 13 1 8.5 22 32 101 0", true, rl.Red, rl.Yellow, 2)
	multi.AddComponent(boom.NewSpriteComp(DOOR))


	// or create a new game object with multiple components
	controllable := boom.Create(
		200, 300, 20, 80,
		boom.NewRectComp(rl.Green, rl.Yellow, 2),
		boom.NewSpriteComp(UGENE),
		boom.NewPolyComp("12 13 1 8.5 22 32 101 10", true, rl.Blue, rl.Purple, 2),
		boom.NewVelocityComp(0.5, 0),
	)

	anim := boom.Sprite(400, 400, UGENE, DOOR)

	anim.Get("sprite").(*boom.SpriteComp).AddAnim("idle", []int{0, 1}, 2, true)
	anim.Get("sprite").(*boom.SpriteComp).Play("idle")

	controllable.SetAngle(12)
	rect.SetAngle(30)
	regpoly.SetScale(2, 3)

	rect.AddComponent(
		boom.NewCollideComp(
			boom.CollisionRect{Width: 50, Height: 50},
			"rect",
		))
	

	control := boom.NewInputComp()
	control.NewInput(rl.KeyLeft, boom.KeyDown, func() {
		control.GameObject.AddAngle(-1)
	})
	control.NewInput(rl.KeyRight, boom.KeyDown, func() {
		control.GameObject.AddAngle(1)
	})

	controllable.AddComponent(control)

	scene.Add(ellipse, rect, sprite, poly, regpoly, multi, controllable, anim)

	game.Run()

}
