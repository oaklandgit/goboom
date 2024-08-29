package main

type Scale struct {
	ScaleX float32
	ScaleY float32
}

func (s *Scale) SetScale(x, y float32) {
	s.ScaleX = x
	s.ScaleY = y
}

func (s *Scale) GetScale() (float32, float32) {
	return s.ScaleX, s.ScaleY
}

func (s *Scale) GetScaleX() float32 {
	return s.ScaleX
}

func (s *Scale) GetScaleY() float32 {
	return s.ScaleY
}