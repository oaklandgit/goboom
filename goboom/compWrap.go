package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type WrapComp struct {
	CompSetGet
	WrapX, WrapY      	bool
	PaddingX, PaddingY  float32
}

func NewWrapComp(wrap ...bool) *WrapComp {

	comp := &WrapComp{}

	if len(wrap) == 1 {
		comp.WrapX = wrap[0]
		comp.WrapY = wrap[0]
		return comp
	}

	if len(wrap) == 2 {
		comp.WrapX = wrap[0]
		comp.WrapY = wrap[1]
		return comp
	}

	comp.WrapX = true
	comp.WrapY = true
	return comp

}

func (w *WrapComp) SetPadding(paddingX, paddingY float32) {
	w.PaddingX = paddingX
	w.PaddingY = paddingY
}

func (*WrapComp) GetComponentId() string {
	return "wrap"
}

func (w *WrapComp) OnInit() {}

func (w *WrapComp) OnUpdate(scene *GameObject) {

	handleWrap := func(coordinate, max, padding float32, wrap bool) float32 {
		if wrap {
			return rl.Wrap(coordinate, -padding, max+padding)
		}
		return coordinate
	}

	width := scene.GetWidth()
	height := scene.GetHeight()
	obj := w.GameObject

	obj.SetX(handleWrap(obj.GetX(), width, w.PaddingX, w.WrapX)) 
	obj.SetY(handleWrap(obj.GetY(), height, w.PaddingY, w.WrapY))

}

func (w *WrapComp) OnDraw(scene *GameObject) {}