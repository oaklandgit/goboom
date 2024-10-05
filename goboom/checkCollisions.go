package goboom

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)


func IsColliding(this, that *CollideComp) bool {

	// object position
	thisPos := rl.NewVector2(
		this.GameObject.GetGlobalX(),
		this.GameObject.GetGlobalY() ,
	)
	thatPos := rl.NewVector2(
		that.GameObject.GetGlobalX(),
		that.GameObject.GetGlobalY(),
	)

	switch thisShape := this.Shape.(type) {
	case CollisionCircle:
		switch thatShape := that.Shape.(type) {
		case CollisionCircle:
			return rl.CheckCollisionCircles(
				rl.Vector2Add(thisPos, thisShape.GetCenter()), thisShape.GetRadius(),
				rl.Vector2Add(thatPos, thatShape.GetCenter()), thatShape.GetRadius())
		case CollisionRect:
			return rl.CheckCollisionCircleRec(
				rl.Vector2Add(thisPos, thisShape.GetCenter()), thisShape.GetRadius(),
				rl.Rectangle{
					X: thatPos.X,
					Y: thatPos.Y,
					Width: thatShape.Width,
					Height: thatShape.Height,
				})		
		}
	case CollisionRect:
		switch thatShape := that.Shape.(type) {
		case CollisionCircle:
			return rl.CheckCollisionCircleRec(
				rl.Vector2Add(thatPos, thatShape.GetCenter()), thatShape.GetRadius(),
				rl.Rectangle{
					X: thisPos.X,
					Y: thisPos.Y,
					Width: thisShape.Width,
					Height: thisShape.Height,
				})
		case CollisionRect:
			return rl.CheckCollisionRecs(
				rl.Rectangle{
					X: thisPos.X,
					Y: thisPos.Y,
					Width: thisShape.Width,
					Height: thisShape.Height,
				},
				rl.Rectangle{
					X: thatPos.X,
					Y: thatPos.Y,
					Width: thatShape.Width,
					Height: thatShape.Height,
				})
		}
	}

	return false

}

func CheckCollisions(scene *GameObject) {

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
				if IsColliding(c1, c2) {

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

					// fmt.Println("COLLISION!", o1.GetId(), "and", o2.GetId())
					c1.CollidingWith = append(c1.CollidingWith, o2)
					c2.CollidingWith = append(c2.CollidingWith, o1)

					// check for colliders
					for _, c := range c1.Colliders {
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