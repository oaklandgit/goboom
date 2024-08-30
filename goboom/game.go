package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Title string
	Width float32
	Height float32
	FPS int
	BgColor rl.Color
	GameObjects []Renderable
	InputManager
}


func NewGame(width, height float32, title string) *Game {
	return &Game{
		Title: title,
		Width: width,
		Height: height,
		BgColor: rl.White,
		FPS: 60,
	}
}

func (g *Game) Add(objects ...Renderable) {
	g.GameObjects = append(g.GameObjects, objects...)
}

func (g *Game) Remove(obj Renderable) {
	for i, o := range g.GameObjects {
		if o == obj {
			g.GameObjects = append(g.GameObjects[:i], g.GameObjects[i+1:]...)
			break
		}
	}
}

func (g *Game) SetBgColor(color rl.Color) {
	g.BgColor = color
}

func (g *Game) Run() {

	debugMode := 0

	DebugOff := func() {}
	DrawGrid := func() {
		DrawGrid(int32(g.Width), int32(g.Height), 12)
		DrawMouseCoordinates()
	}
	var DebugModes = []func(){
		DebugOff,
		DrawGrid,
		DrawPerformance,
	}

	

	g.NewInput(rl.KeyD, KeyPressed, func() {
		debugMode = (debugMode + 1) % len(DebugModes)
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
		g.Update(rl.GetFrameTime())
		g.Draw()
		DebugModes[debugMode]()
		rl.EndDrawing()
	}
	rl.CloseWindow()

}