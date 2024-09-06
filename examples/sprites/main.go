package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const IMAGE_PATH = "../../assets/ugene.png"

func main() {

	game := boom.NewGame(800, 600, "Sprites")
	game.SetBgColor(rl.Black)
	scene := game.GetCurrentScene()

	player := boom.NewSprite(scene.CenterX(), scene.CenterY(), IMAGE_PATH)
	player.SetOrigin(0.5, 0.5)
	player.SetScale(5, 5)

	scene.Add(player)

	game.Run()

}
