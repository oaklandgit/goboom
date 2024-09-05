package goboom

type Shape struct {
	Visible bool
	Opacity float32

	OnDraw func()
	GetWidth func() float32
	GetHeight func() float32
}

func (s *Shape) SetVisible(visible bool) {
	s.Visible = visible
}

func (s *Shape) IsVisible() bool {
	return s.Visible
}
