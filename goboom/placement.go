package goboom

func PutTop(a *GameObject, b *GameObject, offsetX, offsetY float32) {
	centerX := float32(a.GetWidth() / 2)
	b.SetXY(centerX - a.GetWidth()/2 + offsetX, offsetY)
}

func PutLeft(a *GameObject, b *GameObject, offsetX, offsetY float32) {
	centerY := float32(a.GetHeight() / 2)
	b.SetXY(offsetX, centerY - b.GetHeight()/2 + offsetY)
}

func PutCenter(a *GameObject, b *GameObject, offsetX, offsetY float32) {
	centerX := float32(a.GetWidth() / 2)
	centerY := float32(a.GetHeight() / 2)
	b.SetXY(centerX + offsetX - b.GetWidth()/2, centerY + offsetY - b.GetHeight())
}

func PutBottom(a *GameObject, b *GameObject, offsetX, offsetY float32) {
	centerX := float32(a.GetWidth() / 2)
	centerY := float32(a.GetHeight())
	b.SetXY(centerX - b.GetWidth()/2 + offsetX, centerY - b.GetHeight() + offsetY)
}