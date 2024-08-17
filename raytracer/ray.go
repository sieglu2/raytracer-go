package raytracer

import "myhappyland3113.com/raytracer/math3d"

type Ray struct {
	Origin          *math3d.Vector3
	NormalDirection *math3d.Vector3
}

func NewRay(origin, direction *math3d.Vector3) *Ray {
	return &Ray{
		Origin:          origin,
		NormalDirection: direction.Normalize(),
	}
}

func (m *Ray) GetSpecificPoint(distance float64) *math3d.Vector3 {
	return m.Origin.Add(math3d.MulVec3Float64(m.NormalDirection, distance))
}
