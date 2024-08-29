package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func (g *Game) Add(objects ...Lifecycle) {
	g.GameObjects = append(g.GameObjects, objects...)
}


// func (p *GameObject) Add(children ...*GameObject) *GameObject {
// 	for _, c := range children {
// 		p.children = append(p.children, c)
// 		c.parent = p
// 		c.game = p.game
// 	}

// 	return p
// }

func (g *Game) SetBgColor(color rl.Color) {
	g.BgColor = color
}

func (g *Game) Run() {

	rl.InitWindow(int32(g.Width), int32(g.Height), g.Title)
	rl.SetTargetFPS(int32(g.FPS))

	// INITIALIZATION
	// for _, obj := range g.GameObjects {
	// 	obj.OnInit()
	// }

	for !rl.WindowShouldClose() {	
		rl.BeginDrawing()
		rl.ClearBackground(g.BgColor)
		g.Draw()
		rl.EndDrawing()
	}
	rl.CloseWindow()

}

func (g *Game) Draw() {

	for _, obj := range g.GameObjects {
		
		rl.PushMatrix()
		rl.Translatef(
			obj.GetX() + (obj.GetWidth() * obj.GetOriginX()),
			obj.GetY() + (obj.GetHeight() * obj.GetOriginY()), 0)
		rl.Rotatef(obj.GetAngle(), 0, 0, 1)
		rl.Scalef(obj.GetScaleX(), obj.GetScaleY(), 1)
		obj.OnDraw()
		rl.PopMatrix()

	}

}