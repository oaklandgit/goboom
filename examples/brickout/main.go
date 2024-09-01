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

	createBrick := func() boom.Renderable {
		brick := boom.NewRectangle(0, 0, 40, 20, rl.White)
		brick.SetFill(rl.Red)
		return brick
	}

	GridArray(game, 8, 22, 8, 3, 8, createBrick)
	
	game.Run()

}

func GridArray(g *boom.Game, x, y, gap float32, rows, cols int, objFunc func() boom.Renderable) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			obj := objFunc()
			obj.SetXY(
				float32(j * ( int(obj.GetWidth() + gap ))),
				float32(i * ( int(obj.GetHeight() + gap ))))
			obj.SetOrigin(0.5, 0.5)
			g.Add(obj)
		}
	}
}

