package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	// Setup game and scene
	game := boom.NewGame(800, 600, "Mousle Command")
	game.SetBgColor(rl.Black)
	scene := game.GetCurrentScene()

	// Intro text
	intro := boom.Text(0, 0, "Defend Your City!", 30, rl.White)
	intro.SetOrigin(-0.5, -0.5)
	intro.SetLifespan(4_000)

	// Scoreboard
	scoreboard, incrementScore := createScoreboard()

	// Earth
	earth := createEarth(scene)

	// Bases
	bases := boom.GridArray(1, 4, 120, NewBase)

	// Crosshairs
	player := createCrosshairs(scene, incrementScore)

	// Add to scene
	scene.Add(intro, scoreboard, player, earth, bases)
	boom.PutCenter(scene, player, 0, 0)
	boom.PutCenter(scene, bases, -232, 166) // not sure why this won't center
	boom.PutTopCenter(scene, scoreboard, 0, 32)
	boom.PutCenter(scene, intro, 0, 0)

	// Rain down hellfire
	go spawnMissiles(scene)

	game.Run()

}



