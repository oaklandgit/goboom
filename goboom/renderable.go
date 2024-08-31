package goboom

type Renderable interface {
	OnInit()
	OnUpdate()
	OnDraw()

	GetX() float32
	GetY() float32
	GetXY() (float32, float32)
	GetWidth() float32
	GetHeight() float32
	GetScaleX() float32
	GetScaleY() float32
	GetAngle() float32
	GetOriginX() float32
	GetOriginY() float32

	
	SetXY(x, y float32)


	IsVisible() bool
	IsDeleted() bool

	RespondToInputs()
}