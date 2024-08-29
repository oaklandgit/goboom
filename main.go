package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	game := NewGame(800, 600, "My Game")
	game.SetBgColor(rl.Black)

	circle := NewCircle(200, 200, 25, rl.Red)
	rect := NewRectangle(0, 0, 50, 50, rl.Red)
	text := NewText(400, 400, "Hello, World!", rl.Red)
	tri := NewRegPoly(0, 0, 5, 25, rl.Red)
	pent := NewRegPoly(0, 0, 5, 25, rl.Red)

	game.PutCenter(circle, 0, 0)
	game.PutTopLeft(text, 0, 0)
	game.PutTopRight(tri, 0, 0)
	game.PutBottomLeft(rect, 0, 0)
	game.PutBottomRight(pent, 0, 0)

	game.Add(text, circle, rect, tri, pent)

	game.Run()

}