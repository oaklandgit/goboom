package goboom

// type Sprite struct {
// 	Frames []rl.Texture2D
// 	CurrentFrame int
// }

// func NewSprite(x, y float32, path string) *GameObject {

// 	tx := rl.Texture2D{}

// 	s := NewGameObject()
// 	s.X = x
// 	s.Y = y

// 	// must run after the window is initialized
// 	s.OnInit = func () {
// 		tx = rl.LoadTexture(path)
// 	}

// 	s.OnDraw = func () {
// 		rl.DrawTexture(tx, int32(s.X), int32(s.Y), rl.White)
// 	}

// 	s.GetWidth = func() float32 {
// 		return float32(tx.Width)
// 	}

// 	s.GetHeight = func() float32 {
// 		return float32(tx.Height)
// 	}

// 	return s
// }