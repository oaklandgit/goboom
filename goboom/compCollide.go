package goboom

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollideComp struct {
	GameObject 			*GameObject // my game object
	Shape 				ColliderShape // my shape for collisions
	Colliders 			[]Collider // who to collide with and what to do
	CollidingWith 		[]*GameObject // who am I currently colliding with?
	LastCollisionTime 	map[*GameObject]time.Time 
}

type ColliderShape interface {
	GetBoundingBox() rl.Rectangle
	// IsCollidingWith(other *CollideComp) bool
	GetType() string
}

// CIRCLE COLLISION

type CollisionCircle struct {
	Radius float32
}

func (c CollisionCircle) GetBoundingBox() rl.Rectangle {
	r := c.Radius
	return rl.NewRectangle(-r, -r, r*2, r*2)
}

func (c CollisionCircle) GetCenter() rl.Vector2 {
	return rl.NewVector2(0, 0)
}

func (c CollisionCircle) GetRadius() float32 {
	return c.Radius
}

func (c CollisionCircle) GetType() string {
	return "circle"
}

func (c CollisionRect) GetType() string {
	return "rect"
}

func (me CollideComp) IsCollidingWith(other *CollideComp) bool {

	// object position
	mePos := rl.NewVector2(me.GameObject.X, me.GameObject.Y)
	otherPos := rl.NewVector2(other.GameObject.X, other.GameObject.Y)

	// circle vs circle
	if me.Shape.GetType() == "circle" && other.Shape.GetType() == "circle" {
		circle1 := me.Shape.(CollisionCircle)
		circle2 := other.Shape.(CollisionCircle)
		return rl.CheckCollisionCircles(
			rl.Vector2Add(mePos, circle1.GetCenter()), circle1.GetRadius(), 
			rl.Vector2Add(otherPos, circle2.GetCenter()), circle2.GetRadius())
	} 

	// circle vs rect
	if me.Shape.GetType() == "circle" && other.Shape.GetType() == "rect" {
		circle := me.Shape.(CollisionCircle)
		rect := other.Shape.(CollisionRect)
		return rl.CheckCollisionCircleRec(
			rl.Vector2Add(mePos, circle.GetCenter()), circle.GetRadius(),
			rl.Rectangle{
				X: other.GameObject.X,
				Y: other.GameObject.Y,
				Width: rect.Width,
				Height: rect.Height,
			})
	}

	// rect vs circle
	if me.Shape.GetType() == "rect" && other.Shape.GetType() == "circle" {
		rect := me.Shape.(CollisionRect)
		circle := other.Shape.(CollisionCircle)
		return rl.CheckCollisionCircleRec(
			rl.Vector2Add(otherPos, circle.GetCenter()), circle.GetRadius(),
			rl.Rectangle{
				X: me.GameObject.X,
				Y: me.GameObject.Y,
				Width: rect.Width,
				Height: rect.Height,
			})
	}

	// rect vs rect
	if me.Shape.GetType() == "rect" && other.Shape.GetType() == "rect" {
		rect1 := me.Shape.(CollisionRect)
		rect2 := other.Shape.(CollisionRect)
		return rl.CheckCollisionRecs(
			rl.Rectangle{
				X: me.GameObject.X,
				Y: me.GameObject.Y,
				Width: rect1.Width,
				Height: rect1.Height,
			},
			rl.Rectangle{
				X: other.GameObject.X,
				Y: other.GameObject.Y,
				Width: rect2.Width,
				Height: rect2.Height,
			})
	}
	

	return false
}

// RECT COLLISION

type CollisionRect struct {
	Width float32
	Height float32
}

func (c CollisionRect) GetBoundingBox() rl.Rectangle {
	return rl.NewRectangle(0, 0, c.Width, c.Height)
}

// func (c CollisionRect) IsCollidingWith(other *CollideComp) bool {
	
//     if rect, ok := other.Shape.(*CollisionRect); ok {
//         return rl.CheckCollisionRecs(c.GetBoundingBox(), rect.GetBoundingBox())
//     }

//     if circle, ok := other.Shape.(*CollisionCircle); ok {
//         return rl.CheckCollisionCircleRec(circle.GetCenter(), circle.GetRadius(), c.GetBoundingBox())
//     }

//     return false
	
// }


type Collider struct {
	Vs     string
	Action func(obj1, obj2 *GameObject)
}

func NewCollideComp(shape ColliderShape, tags ...string) *CollideComp {
	comp := &CollideComp{
		Shape: shape,
		LastCollisionTime: make(map[*GameObject]time.Time), 
	}
	return comp
}

func (c *CollideComp) NewCollider(vs string, action func(obj1, obj2 *GameObject)) *CollideComp {

	collider := Collider{
		Vs: vs,
		Action: action,
	}

	c.Colliders = append(c.Colliders, collider)

	return c
}

func (c *CollideComp) GetComponentId() string {
	return "collide"
}

func (c *CollideComp) SetGameObject(g *GameObject) {
	c.GameObject = g
}

func (c *CollideComp) GetGameObject() *GameObject {
	return c.GameObject
}


func (c *CollideComp) OnInit() {}

func (c *CollideComp) OnUpdate(scene *GameObject) {

	objs := scene.GetAll()
	collideObjs := []*GameObject{}

	for _, o := range objs {
		if o.HasComponent("collide") {
			collideObjs = append(collideObjs, o)
		}
	}

	// reset collidingWith
	for _, o := range collideObjs {
		o.GetComponent("collide").(*CollideComp).CollidingWith =
			o.GetComponent("collide").(*CollideComp).CollidingWith[:0]
	}

	cooldown := 1 * time.Second
	
	for i, o1 := range collideObjs {
		for j, o2 := range collideObjs {
			if i != j {

				// get the collide components
				c1 := o1.GetComponent("collide").(*CollideComp)
				c2 := o2.GetComponent("collide").(*CollideComp)

				
				// check if they are colliding
				if c1.IsCollidingWith(c2) {

					now := time.Now()

					// Check cooldown for c1
                    if lastTime, ok := c1.LastCollisionTime[o2]; ok {
                        if now.Sub(lastTime) < cooldown {
                            continue // Skip this collision due to cooldown
                        }
                    }

					// Check cooldown for c2
                    if lastTime, ok := c2.LastCollisionTime[o1]; ok {
                        if now.Sub(lastTime) < cooldown {
                            continue // Skip this collision due to cooldown
                        }
                    }

					// Update the last collision time
                    c1.LastCollisionTime[o2] = now
                    c2.LastCollisionTime[o1] = now
					
					// fmt.Println("COLLISION!")
					// fmt.Println("COLLISION!", o1.GetId(), "and", o2.GetId())
					c1.CollidingWith = append(c1.CollidingWith, o2)
					c2.CollidingWith = append(c2.CollidingWith, o1)

					// check for colliders
					for _, c := range c1.Colliders {
						fmt.Println("COLLIDER!")
						if c.Vs == o2.Tags[0] {
							
							c.Action(o1, o2)
						}
					}

					for _, c := range c2.Colliders {
						if c.Vs == o1.Tags[0] {
							c.Action(o2, o1)
						}
					}

				}

			}
		}
	}
}

func (c *CollideComp) OnDraw(scene *GameObject) {

	if c.Shape.GetType() == "circle" {
		circle := c.Shape.(CollisionCircle)
		rl.DrawCircleLines(int32(c.GameObject.X + circle.Radius), int32(c.GameObject.Y + circle.Radius), circle.Radius, rl.Yellow)
		return
	}

	if c.Shape.GetType() == "rect" {
		rect := c.Shape.(CollisionRect)
		// rl.DrawRectangleLines(int32(c.GameObject.X), int32(c.GameObject.Y), int32(rect.Width), int32(rect.Height), rl.Red)
		rl.DrawRectangleLinesEx(rect.GetBoundingBox(), 2, rl.Red)
		return
	}

}

func (c *CollideComp) GetWidth() float32 {
	return 0
}

func (c *CollideComp) GetHeight() float32 {
	return 0
}
