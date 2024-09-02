package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Wrapping struct {
	wrapX    bool
	wrapY    bool
	paddingX float32
	paddingY float32
}

func (o *GameObject) SetWrap(wrap ...bool)  {
	if len(wrap) == 0 {
		o.wrapX = true
		o.wrapY = true
		return
	}

	if len(wrap) == 1 {
		o.wrapX = wrap[0]
		o.wrapY = wrap[0]
		return
	}

	if len(wrap) == 2 {
		o.wrapX = wrap[0]
		o.wrapY = wrap[1]
		return
	}
}

func (o *GameObject) GetWrap() (bool, bool) {
	return o.wrapX, o.wrapY
}

func (o *GameObject) SetPadding(paddingX, paddingY float32) *GameObject {
	o.paddingX = paddingX
	o.paddingY = paddingY
	return o
}

func (o *GameObject) Wrap() {


	handleWrap := func(coordinate, max, padding float32, wrap bool) float32 {
		if wrap {
			return rl.Wrap(coordinate, -padding, max+padding)
		}
		return coordinate
	}

	width := o.GetGame().GetWidth()
	height := o.GetGame().GetHeight()

	o.SetX(handleWrap(o.GetX(), width, o.paddingX, o.wrapX)) 
	o.SetY(handleWrap(o.GetY(), height, o.paddingY, o.wrapY))

}
