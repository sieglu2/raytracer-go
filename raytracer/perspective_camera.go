package raytracer

import (
	"math"

	"myhappyland3113.com/raytracer/math3d"
)

type PerspectiveCamera struct {
	Position *math3d.Vector3
	Front    *math3d.Vector3

	// the actual Up vector passed from user.
	refUp *math3d.Vector3
	right *math3d.Vector3
	// the derived Up vector calculated from Front and refUp.
	derivedUp *math3d.Vector3

	fovHorizontal       float64
	fovVertical         float64
	fovScaledHorizontal float64
	fovScaledVertical   float64
}

func NewPespectiveCamera(position, front, up *math3d.Vector3, fov float64, width, height int) *PerspectiveCamera {
	right := math3d.Cross(front, up)
	derivedUp := math3d.Cross(right, front)

	fovHorizontal := fov
	fovScaledHorizontal := fov * math.Pi / 180.0
	fovVertical := fovHorizontal / (float64(width) / float64(height))
	fovScaledVertical := fovVertical * math.Pi / 180.0

	return &PerspectiveCamera{
		Position: position,
		Front:    front,

		refUp:     up,
		right:     right,
		derivedUp: derivedUp,

		fovHorizontal:       fovHorizontal,
		fovVertical:         fovVertical,
		fovScaledHorizontal: fovScaledHorizontal,
		fovScaledVertical:   fovScaledVertical,
	}
}

func (m *PerspectiveCamera) GenerateRay(x, y float64) *Ray {
	r := math3d.MulVec3Float64(m.right, (x-0.5)*m.fovScaledHorizontal)
	u := math3d.MulVec3Float64(m.derivedUp, (y-0.5)*m.fovScaledVertical)
	return NewRay(m.Position, m.Front.Add(r).Add(u))
}
