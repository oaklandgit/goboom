package main

type Pivot struct {
	OriginX float32
	OriginY float32
}

func (p *Pivot) SetPivot(x, y float32) {
	p.OriginX = x
	p.OriginY = y
}

func (p *Pivot) GetPivot() (float32, float32) {
	return p.OriginX, p.OriginY
}