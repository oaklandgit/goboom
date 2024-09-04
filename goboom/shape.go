package goboom

type Shape struct {
	Width float32
	Height float32	
	OnDraw func()
	GetWidth func() float32
	GetHeight func() float32
}

// GetWidth() float32
// GetHeight() float32