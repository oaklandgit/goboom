package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	Title string
	Width float32
	Height float32
	FPS int
	BgColor rl.Color
	GameObjects []Lifecycle
	
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

func (g *Game) Add(obj ...Lifecycle) {
	g.GameObjects = append(g.GameObjects, obj...)
}

func (g *Game) SetBgColor(color rl.Color) {
	g.BgColor = color
}

func (g *Game) Run() {

	rl.InitWindow(int32(g.Width), int32(g.Height), g.Title)
	rl.SetTargetFPS(int32(g.FPS))

	// INITIALIZATION
	for _, obj := range g.GameObjects {
		obj.OnInit()
	}

	for !rl.WindowShouldClose() {
		
		rl.BeginDrawing()
		rl.ClearBackground(g.BgColor)

		// UPDATE AND DRAW
		for _, obj := range g.GameObjects {
			obj.OnUpdate()
			obj.OnDraw()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()

}