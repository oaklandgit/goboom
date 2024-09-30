package goboom

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawGrid(w, h, every int32, color rl.Color) {
	for x := int32(0); x < w; x += every {
		for y := int32(0); y < h; y += every {
			rl.DrawPixel(x, y, color)
		}
	}
}

func DrawGridSize(x, y, every, fontSize int32, color rl.Color) {
	rl.DrawText(fmt.Sprintf("GRID: %d", every), x, y, fontSize, color)
}

func DrawMouseCoordinates(x, y, fontSize int32, color rl.Color) {

	mouseX := rl.GetMouseX()
	mouseY := rl.GetMouseY()
	text := fmt.Sprintf("X: %d, Y: %d", mouseX, mouseY)
	rl.SetMouseCursor(rl.MouseCursorCrosshair)

	rl.DrawLine(mouseX, 0, mouseX, int32(rl.GetScreenHeight()), color)
	rl.DrawLine(0, mouseY, int32(rl.GetScreenWidth()), mouseY, color)

	rl.DrawText(text, x, y, fontSize, color)
}

func DrawPerformance(x, y, size int32, color rl.Color) {
	perf := fmt.Sprintf("FPS: %d", rl.GetFPS())
	rl.DrawText(perf, x, y, size, color)
}

func DrawObjectCount(x, y, size int32, color rl.Color, objects []*GameObject) {
	count := fmt.Sprintf("OBJECTS: %d", len(objects))
	rl.DrawText(count, x, y, size, color)
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

func DrawColliders(objects []*GameObject, color rl.Color) {
	for _, obj := range objects {
		
		if obj.HasComponent("collide") {

			collide := obj.GetComponent("collide").(*CollideComp)
			shape := collide.Shape
			
			switch shape := shape.(type) {
				case CollisionCircle:
					rl.DrawCircleLines(
						int32(obj.GetX()), int32(obj.GetY()), shape.Radius, color)
				case CollisionRect:
					rl.DrawRectangleLines(
						int32(obj.GetGlobalX()),
						int32(obj.GetGlobalY()),
						int32(shape.Width),
						int32(shape.Height),
						color)
					}

		}
		
	}
		
}