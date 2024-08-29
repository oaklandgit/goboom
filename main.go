package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	game := NewGame(800, 600, "My Game")
	game.SetBgColor(rl.Black)

	circle := NewCircle(100, 100, 50, rl.Red)
	rect := NewRectangle(200, 200, 100, 100, rl.Red)
	text := NewText(200, 200, "Hello, World! You rock!", rl.Red)
	
	// game.Add(circle.Shape, rect.Shape, text.Shape)
	game.Add(text, circle, rect)

	game.Run()

}