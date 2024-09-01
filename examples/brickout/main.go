package main

import (
	boom "goboom/goboom"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

const ROTATE_SPEED = 10


func main() {

	rand.Seed(uint64(time.Now().UnixNano()))
	game := boom.NewGame(800, 600, "Brickout")
	game.SetBgColor(rl.Black)

	scene := game.GetCurrentScene()

	createBrick := func() boom.Renderable {
		brick := boom.NewRectangle(0, 0, 40, 20, rl.White)
		brick.SetFill(rl.Red)
		return brick
	}

	bricks := boom.GridArray(3, 8, 4, createBrick)
	bricks.SetXY(100, 100)

	scene.Add(bricks)
	
	game.Run()

}
