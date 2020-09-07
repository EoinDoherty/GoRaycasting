package main

import (
	"log"
	"math"

	"github.com/EoinDoherty/GoRaycasting/raycasting"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	a := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
	}

	tileSize := 40
	width := len(a) * tileSize * 2
	height := len(a[0]) * tileSize
	rays := raycasting.GenerateFanN(0, math.Pi/2, 40)
	player := raycasting.Player{Rays: rays}

	gm := raycasting.MakeMap(a, tileSize)
	mm := raycasting.InitMiniMap(gm, &player)
	view := raycasting.View{rays, &player}

	// game := raycasting.Game{MiniMap: mm, Player: &player}
	game := raycasting.InitGame(mm, &player, view, width, height)

	ebiten.SetWindowSize(2*len(a)*tileSize, len(a[0])*tileSize)
	ebiten.SetWindowTitle("Raycaster")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatalf("main: %v", err)
	}

	//fmt.Println(math.Mod(0.5, 2*math.Pi))
}
