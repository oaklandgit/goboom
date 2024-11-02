package goboom

import "strings"

// TileDict is a map of rune to GameObject factory functions
type TileDict map[rune]func() *GameObject

// Grid is a 2D slice of GameObject pointers
type Grid [][]*GameObject

// Tilemap is a component that represents a grid of child GameObjects
type Tilemap struct {
	CompSetGet
	TileDict
	Grid
	Width, Height float32
}

type Offset struct {
	X, Y float32
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type DirectionMap map[Direction]Offset

var Directions = DirectionMap{
	Up: Offset{0, -1},
	Down: Offset{0, 1},
	Left: Offset{-1, 0},
	Right: Offset{1, 0},
}

// NewTilemap creates a new GameObject with a Tilemap component
func NewTilemap(dict TileDict, plan string, w, h float32) *GameObject {
	container := NewGameObject()
	tm := &Tilemap{
		TileDict: dict,
		Grid: ObjectsFromPlan(container, w, h, dict, plan),
		Width: w,
		Height: h,
	}

	container.AddComponent(tm)
	return container
}


// SetTileAt sets a GameObject at a specific grid position within the Tilemap
func (tm *Tilemap) SetTile(col, row int, obj *GameObject) {
	if col < 0 || row < 0 || row >= len(tm.Grid) || col >= len(tm.Grid[row]) {
		return
	}
	tm.Grid[row][col] = obj
}

// GetTile gets a GameObject at a specific grid position within the Tilemap
func (tm *Tilemap) GetTile(col, row int) *GameObject {
	if col < 0 || row < 0 || row >= len(tm.Grid) || col >= len(tm.Grid[row]) {
		return nil
	}
	return tm.Grid[row][col]
}

// GetTileAt convenience method to get a tile at a specific global position
func (tm *Tilemap) GetTileAt(gx, gy float32) *GameObject {

	// container global position
	cx, cy := tm.GetGameObject().GetGlobalXY()

	// relative global position
	gx -= cx
	gy -= cy

	col := int(gx / tm.Width)
	row := int(gy / tm.Height)
	return tm.GetTile(col, row)
}

// GetTileTagsAt Convenience method to get tags at a specific global position
func (tm *Tilemap) GetTileTagsAt(gx, gy float32) []string {
	return tm.GetTileAt(gx, gy).GetTags()
}

func (tm *Tilemap) GetTileCoordsAt(gx, gy float32) (int, int) {
	// container global position
	cx, cy := tm.GetGameObject().GetGlobalXY()

	// relative global position
	gx -= cx
	gy -= cy

	col := int(gx / tm.Width)
	row := int(gy / tm.Height)

	return col, row
}

func (tm *Tilemap) GetAdjacentTile(dir Direction) float32 {
	return tm.Width
}

// GetTileTags Convenience method to get tags at a specific grid position
func (tm *Tilemap) GetTileTags(col, row int) []string {
	
	tile := tm.GetTile(col, row)
	if tile == nil {
		return []string{"empty"}
	}
	return tile.GetTags()
}

// GetComponentId Returns the component id, "tilemap"
func (tm *Tilemap) GetComponentId() string {
	return "tilemap"
}

func (tm *Tilemap) OnInit() {}
func (tm *Tilemap) OnUpdate(scene *GameObject) {}
func (tm *Tilemap) OnDraw(scene *GameObject) {}

// ObjectsFromPlan creates a 2D slice of GameObjects from a string plan
// and adds them as children to the containing GameObject
func ObjectsFromPlan(container *GameObject, w, h float32, tiles TileDict, plan string) Grid {

	lines := strings.Split(strings.TrimSpace(plan), "\n")
	grid := make(Grid, len(lines))


	for row, line := range lines {

		codes := []rune(strings.TrimSpace(line))
		grid[row] = make([]*GameObject, 0, len(codes))

		for col, code := range codes {
			tileFunc, exists := tiles[code]
            if !exists {
                continue
            }
			tile := tileFunc()
			if tile == nil {
                continue
            }
			tile.SetXY(float32(col) * w, float32(row) * h)
			grid[row] = append(grid[row], tile)
			container.Add(tile)
		}

	}

	return grid
}
