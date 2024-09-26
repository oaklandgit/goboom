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
	Scenes []*GameObject
	CurrentScene int
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

func (scene *GameObject) Cleanup() {
	for _, c := range scene.GetAll() {
		if c.Delete {
			c.GetParent().Remove(c)
		}
	}
}