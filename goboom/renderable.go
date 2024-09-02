package goboom

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderable interface {

    // Game methods
    SetGame(game *Game)
    GetGame() *Game

    SetParent(parent Renderable)
    GetParent() Renderable
    RemoveParent()

    Add(child ...Renderable)
    GetChildren() []Renderable
    Remove(child Renderable)
    // SetChildren(children []Renderable)

    // Lifecycle methods
    Init()
    Update()
    Draw()
    Wrap()

    SetWrap(wrap ...bool)
    SetPadding(paddingX, paddingY float32)
    GetWrap() (bool, bool)

    // InputHandler methods
    AddInput(key int32, mode ButtonMode, action func())
    CheckInput()


    // Identity methods
    SetId(id string)
    GetId() string
    AddTags(tags ...string)
    GetTags() []string
    RemoveTag(tag string)

    // Position methods
    GetX() float32
    GetY() float32
    GetXY() (float32, float32)
    SetX(x float32)
    SetY(y float32)
    SetXY(x, y float32)

    // Rotation methods
    GetAngle() float32
    SetAngle(angle float32)
    AddAngle(angle float32)

    // BoundingBox methods
    GetBBHeight() float32
    GetBBWidth() float32

    // Size methods
    GetWidth() float32
    GetHeight() float32
    // HalfHeight() float32
    // HalfWidth() float32
    // CenterX() float32
    // CenterY() float32

    // Pivot methods
    GetOrigin() (float32, float32)
    GetOriginX() float32
    GetOriginY() float32
    SetOrigin(x float32, y float32)

    // Scale methods
    GetScale() (float32, float32)
    GetScaleX() float32
    GetScaleY() float32
    SetScale(x float32, y float32)

    // Velocity methods
    GetVel() (float32, float32)
    GetVelX() float32
    GetVelY() float32
    SetVel(x float32, y float32)
    SetVelocityByHeading(heading float32, speed float32)

    // Visibility methods
    IsVisible() bool
    SetVisible(visible bool)

    // Lifespan methods
    IsDeleted() bool
    SetLifespan(millisecs int) *GameObject

    // Stroke methods
    SetStroke(color rl.Color, weight float32) *Stroke
    SetStrokeWeight(weight float32) *Stroke

    // Fill methods
    SetFill(color rl.Color) *Fill

    // Convenience methods
    // PutCenter(a Renderable, b Renderable, offsetX, offsetY float32)
    // PutTop(a *GameObject, b *GameObject, offsetX, offsetY float32)
    // PutBottom(parent Renderable, offsetX, offsetY float32)
    // PutLeft(a *GameObject, b *GameObject, offsetX, offsetY float32)
    // PutRight(parent Renderable, offsetX, offsetY float32)
    // PutTopLeft(parent Renderable, offsetX, offsetY float32)
    // PutTopRight(parent Renderable, offsetX, offsetY float32)
    // PutBottomLeft(parent Renderable, offsetX, offsetY float32)
    // PutBottomRight(parent Renderable, offsetX, offsetY float32)


}