package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"

	"myhappyland3113.com/raytracer/foundation"
	"myhappyland3113.com/raytracer/raytracer"
)

// Define basic types and interfaces for ray tracing.
type Vector struct {
	X, Y, Z float64
}

type Ray struct {
	Origin, Direction Vector
}

type Material struct {
	Color color.RGBA
}

type Sphere struct {
	Center   Vector
	Radius   float64
	Material Material
}

type Camera struct {
	Position Vector
	LookAt   Vector
	FOV      float64
}

// Scene holds all objects to render
type Scene struct {
	Objects []Sphere
	Camera  Camera
}

func (s Sphere) intersect(ray Ray) (bool, float64) {
	// Intersection logic goes here
	return false, 0.0
}

func trace(ray Ray, scene Scene) color.RGBA {
	for _, obj := range scene.Objects {
		if hit, _ := obj.intersect(ray); hit {
			return obj.Material.Color
		}
	}
	return color.RGBA{0, 0, 0, 255} // Return black if no intersection
}

func render(scene Scene) *image.RGBA {
	width := 800
	height := 600
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Render logic here, iterating over all pixels and casting rays
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ray := Ray{
				Origin:    scene.Camera.Position,
				Direction: Vector{X: float64(x), Y: float64(y), Z: 0},
			}
			col := trace(ray, scene)
			img.Set(x, y, col)
		}
	}
	return img
}

func main() {
	// raytracer := raytracer.NewRayTracer("")
	// scene := Scene{
	// 	Objects: []Sphere{
	// 		{
	// 			Center:   Vector{X: 0, Y: 0, Z: 10},
	// 			Radius:   1,
	// 			Material: Material{Color: color.RGBA{255, 0, 0, 255}},
	// 		},
	// 	},
	// 	Camera: Camera{
	// 		Position: Vector{X: 0, Y: 0, Z: 0},
	// 		LookAt:   Vector{X: 0, Y: 0, Z: 1},
	// 		FOV:      90,
	// 	},
	// }

	// img := render(scene)

	// file, err := os.Create("output.png")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// png.Encode(file, img)

	logger := foundation.Logger()

	configFilePathPtr := flag.String("config-path", "", "the yaml config's filepath")
	flag.Parse()

	configFilePath := *configFilePathPtr

	logger.Infof("config filepath: %s", configFilePath)
	if len(configFilePath) == 0 {
		panic("empty config filepath.")
	}

	config, err := raytracer.FromYaml(configFilePath)
	if err != nil {
		panic(fmt.Sprintf("failed to read config: %v", err))
	}

	logger.Infof("config: %+v", config)
}
