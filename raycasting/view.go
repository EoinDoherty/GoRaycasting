package raycasting

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

type View struct {
	Rays   Rays
	Player *Player
}

func (v *View) Update(screen *ebiten.Image) error {
	return nil
}

func (v *View) Draw(screen *ebiten.Image) {
	po := v.Player.Orientation
	// screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	screen.Fill(color.Black)
	rays := v.Rays.Array
	width, height := screen.Size()
	n := len(rays)
	colWidth := int(float64(width) / float64(n))

	center := height / 2

	for i, ray := range rays {
		rayLen := ray.Length

		co := po - ray.Orientation

		if co < 0 {
			co += 2 * math.Pi
		}

		if co > 2*math.Pi {
			co -= 2 * math.Pi
		}

		rayLen = rayLen * math.Cos(co)

		if rayLen > 0 {
			blockHeight := int(float64(height) / rayLen)
			if blockHeight > height {
				blockHeight = height
			}
			img, _ := ebiten.NewImage(colWidth, blockHeight, ebiten.FilterNearest)
			if ray.VerticalIntersect {
				img.Fill(color.RGBA{0, 0, 0xdd, 0xff})
			} else {
				img.Fill(color.RGBA{0, 0, 0xff, 0xff})
			}

			ops := ebiten.DrawImageOptions{}

			ops.GeoM.Translate(float64(i*colWidth), float64(center-blockHeight/2))

			screen.DrawImage(img, &ops)
		}
	}
}
