package goboom

import "strings"

type TileDict map[rune]func() *GameObject
type Grid [][]*GameObject

type Tilemap struct {
	CompSetGet
	TileDict
	Grid
	Width, Height float32
}

func NewTilemap(dict TileDict, plan string, w, h float32) *GameObject {
	container := NewGameObject()
	comp := &Tilemap{
		TileDict: dict,
		Grid: ObjectsFromPlan(container, w, h, dict, plan),
		Width: w,
		Height: h,
	}

	container.AddComponent(comp)
	return container
}

func (tm *Tilemap) SetTileAt(x, y int, obj *GameObject) {
	if x < 0 || y < 0 || y >= len(tm.Grid) || x >= len(tm.Grid[y]) {
		return
	}
	tm.Grid[y][x] = obj
}


func (tm *Tilemap) GetTileAt(x, y int) *GameObject {
	if x < 0 || y < 0 || y >= len(tm.Grid) || x >= len(tm.Grid[y]) {
		return nil
	}
	return tm.Grid[y][x]
}


func (tm *Tilemap) GetComponentId() string {
	return "tilemap"
}

func (tm *Tilemap) OnInit() {}

func (tm *Tilemap) OnUpdate(scene *GameObject) {}

func (tm *Tilemap) OnDraw(scene *GameObject) {}

func ObjectsFromPlan(container *GameObject, w, h float32, tiles TileDict, plan string) Grid {

	lines := strings.Split(strings.TrimSpace(plan), "\n")
	grid := make(Grid, len(lines))


	for row, line := range lines {

		codes := []rune(strings.TrimSpace(line))
		grid[row] = make([]*GameObject, 0, len(codes))

		for col, code := range codes {
			tileFunc, exists := tiles[code]
            if !exists {
                // Handle missing tile code
                continue
            }
			tile := tileFunc()
			if tile == nil {
                // Handle nil tile
                continue
            }
			tile.SetXY(float32(col) * w, float32(row) * h)
			grid[row] = append(grid[row], tile)
			container.Add(tile)
		}

	}

	return grid
}
