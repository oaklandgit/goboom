package goboom

import "strings"

type TileDefs map[rune]func() *GameObject

var tilemap [][]rune

func NewTileMap(tiles TileDefs, plan string, w, h float32) *GameObject {

	m := NewGroup(0, 0)

	lines := strings.Split(strings.TrimSpace(plan), "\n")

	for _, line := range lines {
		tilemap = append(tilemap, []rune(strings.TrimSpace(line)))
	}

	for i, row := range tilemap {
		for j, tile := range row {
			if tileFunc, ok := tiles[tile]; ok {
				obj := tileFunc()
				obj.SetXY(float32(j)*w, float32(i)*h)
				m.Add(obj)
			}
		}
	}

	return m
}