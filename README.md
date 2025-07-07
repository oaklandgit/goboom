# Goboom

## A game engine in Go and Raylib.

Goboom aims to be a simple game engine for Go.

## Philosophy

Goboom draws its inspiration from the wonderful javascript-based game engine [Kaboom](https://kaboomjs.com/) and its successor, [KAPLAY](https://kaplayjs.com/). These engines favor composition over inheritance, where game elements (players, enemies, projectiles, platforms, etc.) derive their functionality "a la carte" in the form of _components_.

This simple **Hello World** should give you the idea. It launches a game window and crawls a message up the screen.

```go
package main

import (
	boom "goboom/goboom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	// set up a game and scene
	game := boom.NewGame(800, 600, "Hello World!")
	game.SetBgColor(rl.Blue)
	scene := game.GetCurrentScene()
	message := boom.Text(0, 0, "Hello World", 30, rl.White)

	// create a movement component
	// with an angle of 0 (upwards) and a speed of 3
	movement := boom.NewVelocityComp(0, 0)
	movement.SetVelocityByHeading(0, 3)

	// create a wrap component
	// true means the object will wrap horizontally and vertically
	wrap := boom.NewWrapComp(true)

	// add the components to the game object
	message.AddComponents(movement, wrap)

	// add the game object to the scene
	scene.Add(message)

	// put the game object in the center of the screen
	boom.PutCenter(scene, message, 0, 0)

	// go!
	game.Run()

}
```

## Other Examples

### Hellfire

![Hellfire](/screencaps/hellfire.png) [View Code](/examples/hellfire)

### Asteroids

![Hellfire](/screencaps/asteroids.png) [View Code](/examples/asteroids)

### Brickout

![Brickout](/screencaps/brickout.png) [View Code](/examples/brickout)

## Features and Components

Goboom currently supports these features:

- **Game Lifecycle** - all game objects have init, update and draw
- **Scene Graph** - create object hierarchies of parents and children
- **Component System** - Imbue game objects with functionality such as Movement, Rotational Velocity, Collisions, Lifespans, Screen Wrapping, and more.

- **Built-In Debugger** - cycle through useful visualizations

![Hellfire](/screencaps/debug-grid.png) Grid Debugger

- **Game Object Primitives and Sprites** - quickly create elements via vector shapes (rectangle, ellipse, polygon, regular polygon, etc.) or by loading images.

- **Convenience Functions** - for creating grids of objects (e.g. Space Invader enemies)

- **Input Handler** - for applying mouse and keyboard events to object behaviors.

- Lifecycles: Game loop and lifecycle
- Transformations: Scale, rotation and pivot points.
- Scene graph: Create parent-child relationships between game objects with `object.Add()`
- Quickly create groups of game objects with `boom.GridArray()`

- Create simple game elements from shapes such as `boom.Circle()`, `boom.Ellipse()`, `boom.Polygon()`, `boom.RegPolygon()` and `boom.Rectangle()`.

### Components

- Add movement to objects with the `Velocity` component.
- Define collisions with `Collision` component.
- Set objects' screen-wrapping behavior with the `Wrap` component.

- Additional components include `boom.NewRotVelocityComp()`,

### Installation / Getting Started

COMING SOON
