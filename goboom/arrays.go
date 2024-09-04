package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

func GridArray(rows, cols int, gap float32, objFunc func() *GameObject) *GameObject {

	group := NewRectangle(0, 0, 0, 0, rl.Blank)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			obj := objFunc()
			obj.SetXY(
				(float32(j) * (obj.GetWidth()) + ((float32(j) * gap))),
				(float32(i) * (obj.GetHeight()) +  ((float32(i) * gap))))

			group.Add(obj)
		}
	}

	group.Width = (objFunc().GetWidth() + gap) * float32(cols)
	group.Height = (objFunc().GetHeight() + gap) * float32(rows)

	return group
}