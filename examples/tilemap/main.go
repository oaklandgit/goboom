package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	game := boom.NewGame(800, 600, "Tilemap Example")
	scene := game.GetCurrentScene()

	// Define tiles
	tileDefs := make(boom.TileDefs)
	tileDefs['#'] = func() *boom.GameObject {
		return boom.Rectangle(0, 0, 40, 40, rl.Gray, rl.White, 2)
	}
	tileDefs['o'] = func() *boom.GameObject {
		return boom.Ellipse(0, 0, 40, 40, rl.Gray, rl.White, 2)
	}

	plan := `
		###########
		#   o     #
		#      o  #
		#         #
		#   o     #
		# o       #
		#         #
		#         #
		###########
	`

	level1 := boom.NewTileMap(tileDefs, plan, 40, 40)

	scene.Add(level1)
	boom.PutCenter(scene, level1, 0, 0)

	game.Run()

}


