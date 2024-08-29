package main

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

// func DrawObjectCoordinates(objs []*GameObject) {
// 	for _, o := range objs {
// 		text := fmt.Sprintf("X: %d, Y: %d", int(o.GX()), int(o.GY()))
// 		rl.DrawText(text, int32(o.GX()), int32(o.GY())-20, 16, rl.DarkGray)
// 	}
// }

func DrawMouseCoordinates() {
	mouseX := rl.GetMouseX()
	mouseY := rl.GetMouseY()
	text := fmt.Sprintf("X: %d, Y: %d", mouseX, mouseY)
	rl.DrawText(text, mouseX, mouseY+20, 16, rl.White)
}

// func DrawBoundingBoxes(objs []*GameObject) {

// 	// for _, o := range objs {

// 	// 	// bounds := TransformBoundingBox(o, BoundingBox(o.LocalBounds())) // cast to BoundingBox

// 	// 	rl.DrawRectangleLines(
// 	// 		int32(bounds.X), int32(bounds.Y), int32(bounds.Width), int32(bounds.Height),
// 	// 		rl.Pink,
// 	// 	)
// 	// }
			
	

// 	for _, o := range objs {
// 		var color rl.Color
// 		if len(o.collidingWith) > 0 {
// 			color = rl.Red
// 		} else {
// 			color = rl.Green
// 		}

// 		rect := o.GlobalBounds()

// 		rl.DrawRectangleLines(
// 			int32(rect.X), int32(rect.Y), int32(rect.Width), int32(rect.Height),
// 			color,
// 		)
// 	}

// }

func DrawPerformance() {
	perf := fmt.Sprintf("FPS: %d", rl.GetFPS())
	rl.DrawText(perf, 20, 20, 16, rl.White)
}

// func DrawObjectCount(objs []*GameObject) {
// 	count := fmt.Sprintf("OBJECTS: %d", len(objs))
// 	rl.DrawText(count, 20, 20, 16, rl.White)
// }
