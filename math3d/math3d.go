package math3d

import (
	"github.com/ungerik/go3d/float64/vec3"
	"github.com/ungerik/go3d/mat4"
)

type Vector3 = vec3.T
type Matrix4x4 = mat4.T

var (
	Vector3_Zero = Vector3(vec3.Zero)
)

func AddVec3Float64(v *Vector3, f float64) *Vector3 {
	v[0] += f
	v[1] += f
	v[2] += f
	return v
}

func MulVec3Float64(v *Vector3, f float64) *Vector3 {
	v[0] *= f
	v[1] *= f
	v[2] *= f
	return v
}

func Cross(a, b *Vector3) *Vector3 {
	res := vec3.Cross(a, b)
	return &res
}
