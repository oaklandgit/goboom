package goboom

func GridArray(rows, cols int, gap float32, objFunc func() *GameObject) *GameObject {

	group := NewGroup(0, 0)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			obj := objFunc()
			// fmt.Println(obj.GetBoundingBox().Width, obj.GetBoundingBox().Height)
			obj.SetXY(
				float32(j) * (obj.GetBoundingBox().Width) + ((float32(j) * gap)),
				float32(i) * (obj.GetBoundingBox().Height) +  ((float32(i) * gap)))

			group.Add(obj)
		}
	}

	return group
}