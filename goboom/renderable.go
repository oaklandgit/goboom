package goboom

type Renderable interface {

	// Lifecycle methods
	OnInit()
	OnUpdate()
	OnDraw()
	OnInput()

	// getters
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
	
	// setters
	SetXY(x, y float32)

	// bools
	IsVisible() bool
	IsDeleted() bool

}