package goboom

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Run() {

	debugMode := 0
	gridSize := int32(10)

	DebugModes := []func() {
		func() {}, // DebugOff
		func() {
			DrawGrid(int32(g.Width), int32(g.Height), gridSize, rl.Brown)
			DrawGridSize(22, 22, gridSize, 20, rl.Yellow)
			DrawMouseCoordinates(22, 46, 20, rl.Yellow)
			
		},
		func() {
			DrawBoundingBoxes(
				g.GetCurrentScene().
				GetChildren(),
				rl.Yellow)
		},
		func() {
			DrawPerformance(22, 22, 20, rl.Yellow)
			DrawObjectCount(22, 46, 20, rl.Yellow, g.GetCurrentScene().GetAll())
		},
	}

	g.AddInput(rl.KeyD, KeyPressed, func() {
		debugMode = (debugMode + 1) % len(DebugModes)
	})

	g.AddInput(rl.KeyLeftBracket, KeyPressed, func() {
		gridSize = int32(math.Max(float64(gridSize-1), 4))
	})

	g.AddInput(rl.KeyRightBracket, KeyPressed, func() {
		gridSize = int32(math.Min(float64(gridSize+1), 20))
	})

	rl.InitWindow(int32(g.Width), int32(g.Height), g.Title)
	rl.SetTargetFPS(int32(g.FPS))

	// INITIALIZATION
	for _, obj := range g.GetCurrentScene().GetAll() {
		for _, c := range obj.GetComponents() {
			c.OnInit()
		}
	}

	for !rl.WindowShouldClose() {	
		rl.BeginDrawing()
		rl.ClearBackground(g.BgColor)

		for _, obj := range g.GetCurrentScene().GetAll() {
			// obj.OnInit()
			for _, c := range obj.GetComponents() {
				c.OnUpdate()
				c.OnDraw()
			}
		}

		// g.Update() // to do: add delta time
		// g.Draw()
		DebugModes[debugMode]()
		rl.EndDrawing()
	}
	rl.CloseWindow()

}