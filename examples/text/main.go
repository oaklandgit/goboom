package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	game := boom.NewGame(800, 450, "Text Example")
	scene := game.GetCurrentScene()

	text := boom.Text(0, 0, "Hello, World!", 30, rl.Black)
	// text.SetOrigin(0.5, 0.5)
	boom.PutCenter(scene, text, 0, 0)
	// text.SetXY(400, 225)

	// txt := text.GetComponent("text").(*boom.TextComp)

	// width := float32(rl.MeasureText(txt.Text, txt.FontSize))

	// fmt.Println(txt.Text, " is width:", width)

	// centerX := float32(scene.GetWidth() / 2)
	// centerY := float32(scene.GetHeight() / 2)

	// text.SetXY(centerX, centerY)

	scene.Print()

	scene.Add(text)

	game.Run()

}