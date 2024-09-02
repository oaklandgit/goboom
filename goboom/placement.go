package goboom

func PutTop(a Renderable, b Renderable, offsetX, offsetY float32) {
	centerX := float32(a.GetWidth() / 2)
	b.SetXY(centerX - a.GetWidth()/2 + offsetX, offsetY)
}

func PutLeft(a Renderable, b Renderable, offsetX, offsetY float32) {
	centerY := float32(a.GetHeight() / 2)
	b.SetXY(offsetX, centerY - b.GetHeight()/2 + offsetY)
}

func PutCenter(a Renderable, b Renderable, offsetX, offsetY float32) {
	centerX := float32(a.GetWidth() / 2)
	centerY := float32(a.GetHeight() / 2)
	b.SetXY(centerX + offsetX - b.GetWidth()/2, centerY + offsetY - b.GetHeight())
}

func PutBottom(a Renderable, b Renderable, offsetX, offsetY float32) {
	centerX := float32(a.GetWidth() / 2)
	centerY := float32(a.GetHeight())
	b.SetXY(centerX - b.GetWidth()/2 + offsetX, centerY - b.GetHeight() + offsetY)
}



// func (g *GameObject) PutRight(p Renderable, offsetX, offsetY float32) {
// 	centerX := float32(g.GetWidth())
// 	centerY := float32(g.GetHeight() / 2)
// 	p.SetXY(centerX - p.GetWidth() + offsetX, centerY - p.GetHeight()/2 + offsetY)
// }

// func (g *GameObject) PutCenter(p Renderable, offsetX, offsetY float32) {
// 	centerX := float32(g.GetWidth() / 2)
// 	centerY := float32(g.GetHeight() / 2)
// 	p.SetXY(centerX + offsetX - p.GetWidth()/2, centerY + offsetY - p.GetHeight())
// }

// func (g *GameObject) PutTopLeft(p Renderable, offsetX, offsetY float32) {
// 	p.SetXY(offsetX, offsetY)
// }

// func (g *GameObject) PutTopRight(p Renderable, offsetX, offsetY float32) {
// 	p.SetXY(float32(g.GetWidth()) - p.GetWidth() + offsetX, offsetY)
// }

// func (g *GameObject) PutBottomLeft(p Renderable, offsetX, offsetY float32) {
// 	p.SetXY(offsetX, float32(g.GetHeight()) - p.GetHeight() + offsetY)
// }

// func (g *GameObject) PutBottomRight(p Renderable, offsetX, offsetY float32) {
// 	p.SetXY(float32(g.GetWidth()) - p.GetWidth() + offsetX, float32(g.GetHeight()) - p.GetHeight() + offsetY)
// }