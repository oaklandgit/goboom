package goboom

import (
	"fmt"
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

	rl.InitWindow(int32(g.Width), int32(g.Height), g.Title)
	rl.SetTargetFPS(int32(g.FPS))

	scene := g.GetCurrentScene()
	
	modes := NewInputComp()
	modes.NewInput(rl.KeyD, KeyPressed, func() {
		debugMode = (debugMode + 1) % len(DebugModes)
		fmt.Println("Debug mode!")
	})
	modes.NewInput(rl.KeyLeftBracket, KeyPressed, func() {
		gridSize = int32(math.Max(float64(gridSize-1), 4))
	})
	modes.NewInput(rl.KeyRightBracket, KeyPressed, func() {
		gridSize = int32(math.Min(float64(gridSize+1), 20))
	})
	scene.AddComponent(modes)

	// Inialize all objects and components
	for _, obj := range scene.GetAll() {
		for _, c := range obj.GetComponents() {
			c.OnInit()
		}
		
		obj.OnInit()
	}

	// Main game loop
	for !rl.WindowShouldClose() {	
		rl.BeginDrawing()
		rl.ClearBackground(g.BgColor)

		// redeclare scene in case it was changed
		scene = g.GetCurrentScene()
		for _, c := range scene.GetComponents() {
			c.OnUpdate(scene)
			c.OnDraw(scene)
		}

		// run any additional updating
		scene.OnUpdate()
		

		// Update and draw all direct children of the scene
		for _, obj := range scene.GetChildren() {
			updateAndDrawChild(obj, scene)
		}

		// run any additional drawing
		scene.OnDraw()

		// clean up any objects that need to be removed
		scene.Cleanup()

		// debug mode
		DebugModes[debugMode]()
		rl.EndDrawing()
	}
	rl.CloseWindow()

}

func updateAndDrawChild(obj *GameObject, scene *GameObject) {
		
	rl.PushMatrix()
	
	centerX := obj.X + obj.GetBoundingBox().Width * obj.GetOriginX()
	centerY := obj.Y + obj.GetBoundingBox().Height * obj.GetOriginY()
	
	// rl.Translatef(obj.GetX(), obj.GetY(), 0)
	rl.Translatef(centerX, centerY, 0)
	rl.Rotatef(obj.GetAngle(), 0, 0, 1)
	rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)
	
	obj.OnUpdate()
	obj.OnDraw()
	
	// components are not nested further
	for _, c := range obj.GetComponents() {
		c.OnUpdate(scene)
		c.OnDraw(scene)
	}

	// but chidren are drawn recursively
	for _, child := range obj.GetChildren() {
		updateAndDrawChild(child, scene)
	}
	
	rl.Translatef(-centerX, -centerY, 0)
	
	rl.PopMatrix()
}