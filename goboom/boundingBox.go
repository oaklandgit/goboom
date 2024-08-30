package goboom

type BoundingBox struct {
	MinX, MinY, MaxX, MaxY float32
}

func (bb *BoundingBox) GetBBWidth() float32 {
    return bb.MaxX - bb.MinX
}

func (bb *BoundingBox) GetBBHeight() float32 {
    return bb.MaxY - bb.MinY
}