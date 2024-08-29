package main

type Position struct {
	X float32
	Y float32
}

func (p *Position) SetXY(x, y float32) {
	p.X = x
	p.Y = y
}

func (p *Position) GetXY() (float32, float32) {
	return p.X, p.Y
}

func (p *Position) GetX() float32 {
	return p.X
}

func (p *Position) GetY() float32 {
	return p.Y
}