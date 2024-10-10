package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	game := boom.NewGame(800, 600, "Tilemap Example")
	game.SetBgColor(rl.Black)
	scene := game.GetCurrentScene()

	// Define tiles
	tileDefs := make(boom.TileDict)
	tileDefs['#'] = func() *boom.GameObject {
		return boom.Rectangle(0, 0, 40, 40, rl.Gray, rl.Red, 2)
	}
	tileDefs['o'] = func() *boom.GameObject {
		return boom.Ellipse(0, 0, 40, 40, rl.Blue, rl.White, 2)
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

	level1 := boom.NewTilemap(tileDefs, plan, 40, 40)

	scene.Add(level1)
	// boom.PutCenter(scene, level1, 0, 0)

	game.Run()

}


