package goboom

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawGrid(w, h, every int32) {
	for x := int32(0); x < w; x += every {
		for y := int32(0); y < h; y += every {
			rl.DrawPixel(x, y, rl.DarkGray)
		}
	}
}

func DrawMouseCoordinates() {
	mouseX := rl.GetMouseX()
	mouseY := rl.GetMouseY()
	text := fmt.Sprintf("X: %d, Y: %d", mouseX, mouseY)
	rl.DrawText(text, mouseX, mouseY+20, 16, rl.White)
}

func DrawPerformance() {
	perf := fmt.Sprintf("FPS: %d", rl.GetFPS())
	rl.DrawText(perf, 20, 20, 16, rl.White)
}

func DrawBoundingBoxes(objects []*GameObject, color rl.Color) {
	for _, r := range objects {
		rl.DrawRectangleLines(
			int32(r.GetBoundingBox().X) - int32(r.GetBoundingBox().Width)/2,
			int32(r.GetBoundingBox().Y) - int32(r.GetBoundingBox().Height)/2,
			int32(r.GetBoundingBox().Width),
			int32(r.GetBoundingBox().Height),
			color)
	}
}