package raycasting

import "fmt"

type GameMap struct {
	Array    [][]int
	TileSize int
}

func MakeMap(mapArray [][]int, tileSize int) GameMap {
	return GameMap{
		Array:    mapArray,
		TileSize: tileSize,
	}
}

func (g *GameMap) printMap() {
	for _, row := range g.Array {
		fmt.Println(row)
	}
}
