package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	game := boom.NewGame(600, 800, "Puckman")
	game.SetBgColor(rl.Black)
	scene := game.GetCurrentScene()

	var gridStr = `
#############
#...........#				
#.##.#.#.##.# 
#....#.#....#
###.##.##.###
#...........#
#.##.###.##.#
#.##.###.##.# 
#.....C.....#
###.##.##.###
#....#.#....#
#.##.#.#.##.#
#...........#
#############
`

	grid := createLevel(gridStr)

	scene.Add(grid)
	boom.PutCenter(scene, grid, 0, 0)

	game.Run()
	


}
