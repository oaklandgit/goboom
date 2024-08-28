package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	game := NewGame(800, 600, "My Game")
	game.SetBgColor(rl.Black)

	circle := NewCircle(100, 100, 50, rl.White)
	rect := NewRectangle(200, 200, 100, 100, rl.White)
	text := NewText(300, 300, "Hello, World! You rock!", rl.White)
	
	game.Add(circle, rect, text)

	game.Run()

}