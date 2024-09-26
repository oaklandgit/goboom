package goboom

type Rotation struct {
	Angle float32
}

func (r *Rotation) SetAngle(angle float32) {
	r.Angle = angle
}

func (r *Rotation) GetAngle() float32 {
	return r.Angle
}

func (r *Rotation) AddAngle(angle float32) {
	r.Angle += angle
}