package goboom

import (
	"fmt"

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

func (g *Game) GetTitle() string {
	return g.Title
}

func (g *Game) GetWidth() float32 {
	return g.Width
}

func (g *Game) GetHeight() float32 {
	return g.Height
}

func (g *Game) NewScene(id string) *GameObject {
	scene := NewRectangle(0, 0, g.GetWidth(), g.GetHeight(), rl.Red) // for debugging
	scene.SetStroke(rl.Red, 3) // for debugging
	scene.SetGame(g)
	scene.SetId(id)
	g.Scenes = append(g.Scenes, scene)
	
	return scene
}

func (g *Game) GetCurrentScene() *GameObject {
	return g.Scenes[g.CurrentScene]
}

func (g *Game) SetScene(index int) {
    if index >= 0 && index < len(g.Scenes) {
        g.CurrentScene = index
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

	DrawBoundingBoxesWrapper := func() {
		DrawBoundingBoxes(g.GetCurrentScene().GetChildren())
	}

	var DebugModes = []func() {
		DebugOff,
		DrawBoundingBoxesWrapper,
		DrawGrid,
		DrawPerformance,
	}

	g.AddInput(rl.KeyD, KeyPressed, func() {
		fmt.Println("Debug mode")
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
		g.Update() // to do: add delta time
		g.Draw()
		DebugModes[debugMode]()
		rl.EndDrawing()
	}
	rl.CloseWindow()

}