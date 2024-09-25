package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	
	game := boom.NewGame(800, 450, "Lifespan Example")
	scene := game.GetCurrentScene()

	// Create a new game object
	text := boom.Text(0, 0, "Goodbye, World (wait for it...)", 30, rl.Black)

	// create a lifespan component and set it to 3 seconds
	death := boom.NewLifespanComp()
	death.SetLifespan(3_000)
	
	// add the lifespan component to the text object
	text.AddComponent(death)

	// add the text object to the scene
	scene.Add(text)

	// run the game
	game.Run()
}