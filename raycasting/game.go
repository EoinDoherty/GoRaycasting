package raycasting

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	MiniMap MiniMap
	Player  *Player
	View    View
	mapImg  *ebiten.Image
	mapOps  *ebiten.DrawImageOptions
	viewImg *ebiten.Image
	viewOps *ebiten.DrawImageOptions
}

const TwoPi float64 = 2 * math.Pi
const Pi2 float64 = math.Pi / 2
const Pi3 float64 = 3 * math.Pi / 2

func InitGame(mm MiniMap, p *Player, v View, windowWidth, windowHeight int) Game {
	width := windowWidth / 2
	mapImg, _ := ebiten.NewImage(width, windowHeight, ebiten.FilterNearest)
	mapOps := ebiten.DrawImageOptions{}

	viewImg, _ := ebiten.NewImage(width, windowHeight, ebiten.FilterNearest)
	viewOps := ebiten.DrawImageOptions{}
	viewOps.GeoM.Translate(float64(width), 0)

	return Game{
		MiniMap: mm,
		Player:  p,
		View:    v,
		mapImg:  mapImg,
		mapOps:  &mapOps,
		viewImg: viewImg,
		viewOps: &viewOps,
	}
}

func (g *Game) UpdateRays() {
	rays := g.Player.Rays.Array
	px, py := g.Player.X, g.Player.Y
	gameMap := g.MiniMap.GameMap.Array

	for i, ray := range rays {
		// Check Horizontal lines
		var mx, my, dof int
		var rx, ry, ra, xo, yo float64

		ra = ray.Orientation
		aTan := -1 / math.Tan(ra)

		// Looking up
		if ra > math.Pi {
			ry = math.Floor(py) - 0.0001
			rx = (py-ry)*aTan + px
			yo = -1
			xo = -yo * aTan
		}

		//Looking down
		if ra < math.Pi {
			ry = math.Floor(py) + 1
			rx = (py-ry)*aTan + px
			yo = 1
			xo = -yo * aTan
		}

		if ra == 0 || ra == math.Pi {
			rx = px
			ry = py
			dof = 8
		}

		for dof < 8 {
			mx = int(rx)
			my = int(ry)

			if my >= 0 && mx >= 0 && my < 8 && mx < 8 && gameMap[my][mx] == 1 {
				dof = 8
			} else {
				rx += xo
				ry += yo
				dof++
			}
		}

		diffX := rx - px
		diffY := ry - py
		horizLength := math.Sqrt(diffX*diffX + diffY*diffY)

		// Check Vertical lines
		// var mx, my, dof int
		// var rx, ry, ra, xo, yo float64
		dof = 0

		// ra = ray.Orientation
		nTan := -math.Tan(ra)

		// Looking left
		if ra > Pi2 && ra < Pi3 {
			rx = math.Floor(px) - 0.0001
			ry = (px-rx)*nTan + py
			xo = -1
			yo = -xo * nTan
		}

		//Looking right
		if ra < Pi2 || ra > Pi3 {
			rx = math.Floor(px) + 1
			ry = (px-rx)*nTan + py
			xo = 1
			yo = -xo * nTan
		}

		if ra == 0 || ra == math.Pi {
			rx = px
			ry = py
			dof = 8
		}

		for dof < 8 {
			mx = int(rx)
			my = int(ry)

			if my >= 0 && mx >= 0 && my < 8 && mx < 8 && gameMap[my][mx] == 1 {
				dof = 8
			} else {
				rx += xo
				ry += yo
				dof++
			}
		}

		diffX = rx - px
		diffY = ry - py
		vertLength := math.Sqrt(diffX*diffX + diffY*diffY)

		if horizLength < vertLength {
			g.Player.Rays.Array[i].Length = horizLength
			g.Player.Rays.Array[i].VerticalIntersect = false
		} else {
			g.Player.Rays.Array[i].Length = vertLength
			g.Player.Rays.Array[i].VerticalIntersect = true
		}
		// g.Player.Rays.Array[i].Length = math.Min(horizLength, vertLength)
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.UpdateRays()
	return g.MiniMap.Update(screen)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.MiniMap.Draw(g.mapImg)
	g.View.Draw(g.viewImg)

	screen.DrawImage(g.mapImg, g.mapOps)
	screen.DrawImage(g.viewImg, g.viewOps)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
