package main

type Lifecycle interface {
	OnInit()
	OnUpdate()
	OnDraw()

	GetX() float32
	GetY() float32
	GetWidth() float32
	GetHeight() float32
	GetScaleX() float32
	GetScaleY() float32
	GetAngle() float32
	GetOriginX() float32
	GetOriginY() float32

}