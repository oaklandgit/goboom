package goboom

// SCENE GRAPH CURRENTLY WORKS! DON'T BREAK IT!

// func (g *Game) Draw() {

// 	scene := g.GetCurrentScene()
// 	scene.OnDraw()

// 	for _, obj := range scene.GetChildren() {
// 		obj.OnDraw()
//     }
// }

// func drawWithTransforms(obj *GameObject) {

// 		obj.OnDraw()

// 		if len(obj.GetChildren()) > 0 {

// 			rl.Translatef(obj.GetX(), obj.GetY(), 0)

// 			for _, c := range obj.GetChildren() {
// 				drawWithTransforms(c)
// 			}
// 		}

// }