package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	// set up a game and scene
	game := boom.NewGame(800, 600, "Hello World!")
	game.SetBgColor(rl.Blue)
	scene := game.GetCurrentScene()
	message := boom.Text(0, 0, "Hello World", 30, rl.White)

	// create a movement component
	// with an angle of 0 (upwards) and a speed of 3
	movement := boom.NewVelocityComp(0, 0)
	movement.SetVelocityByHeading(0, 3)

	// create a wrap component
	// true means the object will wrap horizontally and vertically
	wrap := boom.NewWrapComp(true)

	// add the components to the game object
	message.AddComponents(movement, wrap)

	// add the game object to the scene
	scene.Add(message)

	// put the game object in the center of the screen
	boom.PutCenter(scene, message, 0, 0)

	// have at it!
	game.Run()

}