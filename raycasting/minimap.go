package raycasting

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type MiniMap struct {
	GameMap       GameMap
	Player        *Player
	tiles         [][]*ebiten.Image
	tileOptions   [][]ebiten.DrawImageOptions
	playerTile    *ebiten.Image
	playerOptions *ebiten.DrawImageOptions
}

func InitMiniMap(gm GameMap, player *Player) MiniMap {
	mm := MiniMap{
		GameMap: gm,
		Player:  player,
	}

	mm.Init()

	return mm
}

func (m *MiniMap) Init() {
	arr := m.GameMap.Array
	ts := m.GameMap.TileSize

	m.tiles = make([][]*ebiten.Image, len(arr))
	m.tileOptions = make([][]ebiten.DrawImageOptions, len(arr))

	for y, row := range arr {
		m.tiles[y] = make([]*ebiten.Image, len(arr[y]))
		m.tileOptions[y] = make([]ebiten.DrawImageOptions, len(arr[y]))

		for x, cellVal := range row {
			img, _ := ebiten.NewImage(ts, ts, ebiten.FilterNearest)
			if cellVal == 1 {
				img.Fill(color.White)
			} else {
				img.Fill(color.Black)
			}

			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate((float64)(x*ts), (float64)(y*ts))

			m.tiles[y][x] = img
			m.tileOptions[y][x] = op
		}
	}

	m.playerTile, _ = ebiten.NewImage(ts/4, ts/4, ebiten.FilterNearest)
	m.playerTile.Fill(color.RGBA{0xff, 0, 0, 0xff})
	m.playerOptions = &ebiten.DrawImageOptions{}
	m.playerOptions.GeoM.Translate(m.Player.X, m.Player.Y)
}

func (m *MiniMap) Update(screen *ebiten.Image) error {
	return m.Player.Update()
}

func (m *MiniMap) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 0xff})

	images := m.tiles
	options := m.tileOptions

	for y, row := range images {
		for x, img := range row {
			screen.DrawImage(img, &options[y][x])
		}
	}

	tsf := float64(m.GameMap.TileSize)

	// Draw the player
	px := tsf * m.Player.X
	py := tsf * m.Player.Y

	pWidth, pHeight := m.playerTile.Size()
	centerX, centerY := px-float64(pWidth)/2, py-float64(pHeight)/2
	m.playerOptions.GeoM.Reset()
	m.playerOptions.GeoM.Translate(centerX, centerY)
	screen.DrawImage(m.playerTile, m.playerOptions)

	//Draw the rays
	rays := m.Player.Rays.Array

	x1 := px
	y1 := py

	for _, ray := range rays {
		dX := math.Cos(ray.Orientation)
		dY := math.Sin(ray.Orientation)

		scaling := ray.Length * tsf / math.Sqrt(math.Pow(dX, 2)+math.Pow(dY, 2))

		dX *= scaling
		dY *= scaling

		x2, y2 := x1+dX, y1+dY

		ebitenutil.DrawLine(screen, x1, y1, x2, y2, color.RGBA{0, 0xff, 0, 0xff})
	}
}

// func (m *MiniMap) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return outsideWidth, outsideHeight
// }
