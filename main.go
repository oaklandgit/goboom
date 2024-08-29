package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	game := NewGame(800, 600, "My Game")
	game.SetBgColor(rl.Black)

	circle := NewCircle(200, 200, 25, rl.White)
	circle.SetFill(rl.Blue)
	circle.SetStroke(rl.Blank, 2)
	rect := NewRectangle(0, 0, 50, 50, rl.White)
	text := NewText(400, 400, "Hello, World!", rl.White)
	tri := NewRegPoly(0, 0, 5, 25, rl.White)
	pent := NewRegPoly(0, 0, 5, 25, rl.White)
	star := NewPolygon(
		0, 0, rl.White, true,
		0, 50, 14, 15, 47, 15, 23, -7, 30, -42, 0, -25, -30, -42, -23, -7, -47, 15, -14, 15)

	rect2 := NewRectangle(0, 0, 20, 40, rl.White)
	rect2.SetOrigin(0.5, 0.5)

	game.NewInput(rl.KeyRight, KeyDown, func() {
		rect2.SetAngle(rect2.GetAngle() + 10)
	})

	game.NewInput(rl.KeyLeft, KeyDown, func() {
		rect2.SetAngle(rect2.GetAngle() - 10)
	})

	game.NewInput(rl.KeyUp, KeyReleased, func() {
		rect2.SetVelocityByHeading(rect2.GetAngle(), 1)
	})

	game.PutLeft(circle, 0, 0)
	game.PutTopLeft(text, 0, 0)
	game.PutTopRight(tri, 0, 0)
	game.PutBottomLeft(rect, 0, 0)
	game.PutBottomRight(pent, 0, 0)
	game.PutCenter(rect2, 0, 0)
	game.PutCenter(star, 0, 0)

	tri.SetVel(-1, 1)
	star.SetVelocityByHeading(45, 0.1)
	star.SetAngle(30)

	game.Add(text, circle, rect, tri, pent, rect2, star)

	game.Run()

}