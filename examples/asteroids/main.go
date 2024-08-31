package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ROTATE_SPEED = 10

func main() {

	game := boom.NewGame(800, 600, "Asteroids")
	game.SetBgColor(rl.Black)

	ship := createPlayer(game)

	game.PutCenter(ship, 0, 0)
	game.Add(ship)

	game.Run()

}