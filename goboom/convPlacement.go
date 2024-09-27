package goboom

func PutTop(a *GameObject, b *GameObject, offsetX, offsetY float32) {
	centerX := float32(a.GetWidth() / 2)
	b.SetXY(centerX - a.GetWidth()/2 + offsetX, offsetY)
}

func PutTopCenter(a *GameObject, b *GameObject, offsetX, offsetY float32) {
	b.OnInit = func() {
		centerX := float32(a.GetWidth() / 2)
		b.SetXY(centerX - b.GetWidth()/2 + offsetX, offsetY)
	}
}

func PutLeft(a *GameObject, b *GameObject, offsetX, offsetY float32) {
	centerY := float32(a.GetHeight() / 2)
	b.SetXY(offsetX, centerY - b.GetHeight()/2 + offsetY)
}

func PutCenter(a *GameObject, b *GameObject, offsetX, offsetY float32) {

	b.OnInit = func() {

		centerX := float32(a.GetBoundingBox().Width/2)
		centerY := float32(a.GetBoundingBox().Height/2)
		b.SetXY(
			centerX + offsetX - b.GetBoundingBox().Width/2,
			centerY + offsetY - b.GetBoundingBox().Height/2)

	}
		
}

func PutBottom(a *GameObject, b *GameObject, offsetX, offsetY float32) {
	centerX := float32(a.GetWidth() / 2)
	centerY := float32(a.GetHeight())
	b.SetXY(centerX - b.GetWidth()/2 + offsetX, centerY - b.GetHeight() + offsetY)
}