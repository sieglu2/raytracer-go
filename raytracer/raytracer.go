package raytracer

import "myhappyland3113.com/raytracer/math3d"

type RayTracer struct {
	width, height int

	outputFilepath string

	// camera
	// scene
	// lights
	backgroundColor math3d.Vector3

	maxDepth          int
	maxRecursionDepth int
}

func NewRayTracer(configFilepath string) *RayTracer {
	// rayTracer := &RayTracer{}
	return nil
}

func (r *RayTracer) Destroy() error {
	return nil
}

func (r *RayTracer) Save() error {
	return nil
}

func (r *RayTracer) Render() error {
	return nil
}

func (r *RayTracer) RayTraceRecursively() math3d.Vector3 {
	return math3d.Vector3_Zero
}
