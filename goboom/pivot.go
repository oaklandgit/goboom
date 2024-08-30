package goboom

type Pivot struct {
	OriginX float32
	OriginY float32
}

func (p *Pivot) SetOrigin(x, y float32) {
	p.OriginX = x
	p.OriginY = y
}

func (p *Pivot) GetOrigin() (float32, float32) {
	return p.OriginX, p.OriginY
}

func (p *Pivot) GetOriginX() float32 {
	return p.OriginX
}

func (p *Pivot) GetOriginY() float32 {
	return p.OriginY
}