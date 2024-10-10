package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Run() {

	gridSizes := []int32{8, 16, 32, 64, 128}
	debugColors := []rl.Color{rl.Red, rl.Blue, rl.Green, rl.Orange, rl.Purple}

	debugModes := []func() {
		func() {}, // DebugOff
		func() {
			DrawColliders(g.GetCurrentScene().GetAll(), debugColors[0])
		},
		func() {
			DrawGrid(int32(g.Width), int32(g.Height), gridSizes[0], debugColors[0])
			DrawGridSize(22, 22, gridSizes[0], 20, debugColors[0])
			DrawMouseCoordinates(22, 46, 20, debugColors[0])		
		},
		func() {
			DrawBoundingBoxes(
				g.GetCurrentScene().
				GetChildren(),
				debugColors[0])
		},
		func() {
			DrawPerformance(22, 22, 20, debugColors[0])
			DrawObjectCount(22, 46, 20, debugColors[0], g.GetCurrentScene().GetAll())
		},
		func() {
			// DrawColliders(g.GetCurrentScene().GetAll(), debugColors[0])
			DrawGrid(int32(g.Width), int32(g.Height), gridSizes[0], debugColors[0])
			DrawGridSize(22, 22, gridSizes[0], 20, debugColors[0])
			DrawMouseCoordinates(22, 46, 20, debugColors[0])
		},
	}

	rl.InitWindow(int32(g.Width), int32(g.Height), g.Title)
	rl.SetTargetFPS(int32(g.FPS))

	scene := g.GetCurrentScene()
	
	modes := NewInputComp()
	modes.NewInput(rl.KeyD, KeyPressed, func() {
		debugModes = append(debugModes[1:], debugModes[0])
	})
	modes.NewInput(rl.KeyG, KeyPressed, func() {
		gridSizes = append(gridSizes[1:], gridSizes[0])
	})
	modes.NewInput(rl.KeyC, KeyPressed, func() {
		debugColors = append(debugColors[1:], debugColors[0])

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
		CheckCollisions(scene)
		

		// Update and draw all direct children of the scene
		for _, obj := range scene.GetChildren() {
			updateAndDrawChild(obj, scene)
		}

		scene.OnDraw()
		scene.Cleanup()
		debugModes[0]()

		rl.EndDrawing()
	}
	rl.CloseWindow()

}

func updateAndDrawChild(obj *GameObject, scene *GameObject) {

		
	rl.PushMatrix()
	
	centerX := obj.X + obj.GetBoundingBox().Width * obj.GetOriginX()
	centerY := obj.Y + obj.GetBoundingBox().Height * obj.GetOriginY()
	// centerX := obj.GetGlobalX() + obj.GetBoundingBox().Width * obj.GetOriginX()
	// centerY := obj.GetGlobalY() + obj.GetBoundingBox().Height * obj.GetOriginY()
	
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