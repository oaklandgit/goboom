package goboom

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Title string
	Width float32
	Height float32
	FPS int
	BgColor rl.Color
	Scenes []*GameObject
	CurrentScene int
	InputHandler
}


func NewGame(width, height float32, title string) *Game {

	g := &Game{
		Title: title,
		Width: width,
		Height: height,
		BgColor: rl.White,
		FPS: 60,
		Scenes: []*GameObject{},
		CurrentScene: 0,
	}

	g.NewScene("default")
	return g
}

func (g *Game) SetBgColor(color rl.Color) {
	g.BgColor = color
}

func (g *Game) GetTitle() string {
	return g.Title
}

func (g *Game) GetWidth() float32 {
	return g.Width
}

func (g *Game) GetHeight() float32 {
	return g.Height
}

func (g *Game) Update() {

	g.CheckInput()

	scene := g.GetCurrentScene()
	scene.OnUpdate()
	scene.CheckCollisions()

	for _, obj := range scene.GetChildren() {

		if obj.IsDeleted() {
			scene.Remove(obj)
			continue
		}

		obj.OnUpdate()
		obj.OnWrap()
		obj.CheckInput()
	}

}

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
	// for _, obj := range g.GameObjects {
	// 	obj.OnInit()
	// }

	for !rl.WindowShouldClose() {	
		rl.BeginDrawing()
		rl.ClearBackground(g.BgColor)
		g.Update() // to do: add delta time
		g.Draw()
		DebugModes[debugMode]()
		rl.EndDrawing()
	}
	rl.CloseWindow()

}