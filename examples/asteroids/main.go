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

	// ship.SetVel(10, 10)

	game.PutCenter(ship, 0, 0)
	game.Add(ship)

	game.Run()

}