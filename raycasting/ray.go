package raycasting

import (
	"math"
)

type Ray struct {
	Orientation       float64
	Length            float64
	VerticalIntersect bool
}

type Rays struct {
	Array []Ray
}

const TWO_PI float64 = math.Pi * 2

func (r Rays) Rotate(rads float64) {
	for i := range r.Array {
		o := r.Array[i].Orientation
		o += rads

		if o < 0 {
			o += TWO_PI
		} else if o > TWO_PI {
			o -= TWO_PI
		}

		r.Array[i].Orientation = o
	}
}

func GenerateFanN(centerOrientation, spanRadians, numRays float64) Rays {
	return GenerateFan(centerOrientation, spanRadians, spanRadians/numRays)
}

func GenerateFan(centerOrientation float64, spanRadians float64, stepSize float64) Rays {
	halfSpan := spanRadians / 2
	start := centerOrientation - halfSpan
	end := centerOrientation + halfSpan

	rays := make([]Ray, 0)

	for o := start; o < end; o += stepSize {
		ray := Ray{
			Orientation: o,
			Length:      5,
		}
		rays = append(rays, ray)
	}

	return Rays{Array: rays}
}
