package raycasting

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	X           float64
	Y           float64
	Orientation float64
	Rays        Rays
}

func scaleVector(x, y, scale float64) (nX, nY float64) {
	mag := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	return scale * x / mag, scale * y / mag
}

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		dX := math.Cos(p.Orientation)
		dY := math.Sin(p.Orientation)
		dX, dY = scaleVector(dX, dY, 0.05)

		p.X += dX
		p.Y += dY
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		dX := math.Cos(p.Orientation)
		dY := math.Sin(p.Orientation)
		dX, dY = scaleVector(dX, dY, 0.05)

		p.X -= dX
		p.Y -= dY
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Orientation -= 0.1
		p.Rays.Rotate(-0.1)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Orientation += 0.1
		p.Rays.Rotate(0.1)
	}

	if p.Orientation < 0 {
		p.Orientation += TwoPi
	} else if p.Orientation > TwoPi {
		p.Orientation -= TwoPi
	}

	// fmt.Println(p.Orientation)

	return nil
}

// func (p *Player) Draw(screen *ebiten.Image) {
// 	img, _ := ebiten.NewImage(5, 5, ebiten.FilterNearest)
// 	img.Fill(color.RGBA{0xff, 0, 0, 0xff})
// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(p.X, p.Y)
// 	screen.DrawImage(img, op)

// 	for i := 0.0; i < 10.0; i += 1 {
// 		newOrientation := p.Orientation - 0.1*i
// 		dX := math.Cos(newOrientation)
// 		dY := math.Sin(newOrientation)
// 		dX, dY = normalizeVector(dX, dY)

// 		dX = dX * 100
// 		dY = dY * 100

// 		centerX := p.X + 2.5
// 		centerY := p.Y + 2.5

// 		ebitenutil.DrawLine(screen, centerX, centerY, centerX+dX, centerY+dY, color.RGBA{0, 0xff, 0, 0xff})
// 	}

// 	for i := 0.0; i < 10.0; i += 1 {
// 		newOrientation := p.Orientation + 0.1*i
// 		dX := math.Cos(newOrientation)
// 		dY := math.Sin(newOrientation)
// 		dX, dY = normalizeVector(dX, dY)

// 		dX = dX * 100
// 		dY = dY * 100

// 		centerX := p.X + 2.5
// 		centerY := p.Y + 2.5

// 		ebitenutil.DrawLine(screen, centerX, centerY, centerX+dX, centerY+dY, color.RGBA{0, 0xff, 0, 0xff})
// 	}
// }
