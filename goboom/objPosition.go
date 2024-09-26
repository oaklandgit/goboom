package goboom

type Position struct {
	X float32
	Y float32
}

func (p *Position) SetX(x float32) {
	p.X = x
}

func (p *Position) SetY(y float32) {
	p.Y = y
}

func (p *Position) SetXY(x, y float32) {
	p.SetX(x)
	p.SetY(y)
}

func (p *Position) GetX() float32 {
	return p.X
}

func (p *Position) GetY() float32 {
	return p.Y
}

func (p *Position) GetXY() (float32, float32) {
	return p.GetX(), p.GetY()
}