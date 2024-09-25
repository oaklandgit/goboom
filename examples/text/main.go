package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	game := boom.NewGame(800, 450, "Text Example")
	scene := game.GetCurrentScene()

	text := boom.Text(0, 0, "Hello, World!", 30, rl.Black)
	text.SetOrigin(0.5, 0.5)
	boom.PutCenter(scene, text, 0, 0)

	scene.Print()

	scene.Add(text)

	game.Run()

}