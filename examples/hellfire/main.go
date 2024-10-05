package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	// Setup game and scene
	game := boom.NewGame(800, 600, "Hellfire")
	game.SetBgColor(rl.Black)
	scene := game.GetCurrentScene()

	// Intro text
	intro := boom.Text(0, 0, "Defend Your City!", 30, rl.White)
	intro.SetLifespan(4_000)

	// Scoreboard
	scoreboard, incrementScore := createScoreboard()

	// Earth
	earth := createEarth(scene)

	// Bases
	bases := boom.GridArray(1, 4, 120, NewBase)

	// Crosshairs
	player := createCrosshairs(scene, incrementScore)

	// Position everything
	boom.PutCenter(scene, intro, 0, 0)
	boom.PutCenter(scene, player, 0, 0)
	boom.PutCenter(scene, bases, 0, 180)
	boom.PutTopCenter(scene, scoreboard, 0, 32)
	boom.PutBottom(scene, earth, 0, 0)

	// Add everything to the scene
	scene.Add(earth, bases, intro, scoreboard, player)

	// Rain down hellfire
	go spawnMissiles(scene)

	game.Run()

}



