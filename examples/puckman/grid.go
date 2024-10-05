package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CELL_SIZE = 22
	DOT_SIZE = 3
	PLAYER_SPEED = 4
)

var tileDefs = boom.TileDefs{
	'#': func() *boom.GameObject {
		wall := boom.Rectangle(0, 0, CELL_SIZE, CELL_SIZE, rl.Blue, rl.Black, 2)
		collide := boom.NewCollideComp(boom.CollisionRect{Width: CELL_SIZE, Height: CELL_SIZE})
		wall.AddComponents(collide)
		wall.AddTags("wall")
		return wall
	},
	'.': func() *boom.GameObject {
		cell := boom.Rectangle(0, 0, CELL_SIZE, CELL_SIZE, rl.Blank, rl.Blank, 0)
		dot := boom.Circle(CELL_SIZE/2 - DOT_SIZE, CELL_SIZE/2 - DOT_SIZE, DOT_SIZE, rl.White, rl.Black, 2)
		cell.AddTags("dot")
		cell.Add(dot)
		return cell
	},
	'C': func() *boom.GameObject {
		player := boom.Circle(0, 0, CELL_SIZE/2, rl.Yellow, rl.Black, 2)
		player.AddTags("player")

		vel := boom.NewVelocityComp(0, 0)
		control := boom.NewInputComp()
		collide := boom.NewCollideComp(boom.CollisionRect{Width: CELL_SIZE - 4, Height: CELL_SIZE - 4})

		collide.NewCollider("wall", func(p, w *boom.GameObject) {
			vel.SetVelocity(0, 0)
		})


		control.NewInput(rl.KeyUp, boom.KeyPressed, func() {
			vel.SetVelocity(rl.NewVector2(0, -PLAYER_SPEED))
		})
		control.NewInput(rl.KeyDown, boom.KeyPressed, func() {
			vel.SetVelocity(rl.NewVector2(0, PLAYER_SPEED))
		})
		control.NewInput(rl.KeyLeft, boom.KeyPressed, func() {
			vel.SetVelocity(rl.NewVector2(-PLAYER_SPEED, 0))
		})
		control.NewInput(rl.KeyRight, boom.KeyPressed, func() {
			vel.SetVelocity(rl.NewVector2(PLAYER_SPEED, 0))
		})
		player.AddComponents(vel, collide, control)
		return player
	},
}

func createLevel(plan string) *boom.GameObject {
	grid := boom.NewTileMap(tileDefs, plan, CELL_SIZE, CELL_SIZE)
	return grid
}