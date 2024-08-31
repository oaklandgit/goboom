package goboom

func (g *Game) PutTop(p Renderable, offsetX, offsetY float32) {
	centerX := float32(g.Width / 2)
	p.SetXY(centerX - p.GetWidth()/2 + offsetX, offsetY)
}

func (g *Game) PutBottom(p Renderable, offsetX, offsetY float32) {
	centerX := float32(g.Width / 2)
	centerY := float32(g.Height)
	p.SetXY(centerX - p.GetWidth()/2 + offsetX, centerY - p.GetHeight() + offsetY)
}

func (g *Game) PutLeft(p Renderable, offsetX, offsetY float32) {
	centerY := float32(g.Height / 2)
	p.SetXY(offsetX, centerY - p.GetHeight()/2 + offsetY)
}

func (g *Game) PutRight(p Renderable, offsetX, offsetY float32) {
	centerX := float32(g.Width)
	centerY := float32(g.Height / 2)
	p.SetXY(centerX - p.GetWidth() + offsetX, centerY - p.GetHeight()/2 + offsetY)
}

func (g *Game) PutCenter(p Renderable, offsetX, offsetY float32) {
	centerX := float32(g.Width / 2)
	centerY := float32(g.Height / 2)
	p.SetXY(centerX + offsetX - p.GetWidth()/2, centerY + offsetY - p.GetHeight())
}

func (g *Game) PutTopLeft(p Renderable, offsetX, offsetY float32) {
	p.SetXY(offsetX, offsetY)
}

func (g *Game) PutTopRight(p Renderable, offsetX, offsetY float32) {
	p.SetXY(float32(g.Width) - p.GetWidth() + offsetX, offsetY)
}

func (g *Game) PutBottomLeft(p Renderable, offsetX, offsetY float32) {
	p.SetXY(offsetX, float32(g.Height) - p.GetHeight() + offsetY)
}

func (g *Game) PutBottomRight(p Renderable, offsetX, offsetY float32) {
	p.SetXY(float32(g.Width) - p.GetWidth() + offsetX, float32(g.Height) - p.GetHeight() + offsetY)
}