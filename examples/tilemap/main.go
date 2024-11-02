package main

import (
	"fmt"
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
		rect := boom.Rectangle(0, 0, 40, 40, rl.Blue, rl.Black, 2)
		rect.AddTags("wall")
		return rect
	}
	tileDefs['o'] = func() *boom.GameObject {
		circle := boom.Ellipse(0, 0, 40, 40, rl.Blue, rl.White, 2)
		circle.AddTags("dot")
		return circle
	}
	tileDefs['·'] = func() *boom.GameObject { return nil }

	plan := `
		###########
		#··o······#
		#····o····#
		#·········#
		#·········#
		#·········#
		#·········#
		#·········#
		###########
	`

	level1 := boom.NewTilemap(tileDefs, plan, 40, 40)

	scene.Add(level1)
	boom.PutCenter(scene, level1, 0, 0)

	// add a mouse hover
	crosshairs := boom.NewGameObject()
	mouse := boom.NewMouseComp()
	mouse.OnMove(func(x, y float32) {
		grid := level1.GetComponent("tilemap").(*boom.Tilemap)
		tileX, tileY := grid.GetTileCoordsAt(x, y)
		fmt.Println(tileX, tileY)
		tags := grid.GetTileTags(tileX, tileY)
		fmt.Println(tags)
	})
	crosshairs.AddComponent(mouse)
	scene.Add(crosshairs)

	game.Run()

}


