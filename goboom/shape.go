package goboom

type Shape struct {
	OnDraw func()
	GetWidth func() float32
	GetHeight func() float32
	Opacity float32
}
