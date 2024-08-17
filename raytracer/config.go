package raytracer

import (
	"os"

	"gopkg.in/yaml.v2"
)

type MaterialType string
type LightType string

const (
	MaterialType_Ambient  = MaterialType("ambient")
	MaterialType_Diffuse  = MaterialType("diffuse")
	MaterialType_Specular = MaterialType("specular")

	LightType_PointLight     = LightType("point_light")
	LightType_DirectionLight = LightType("direction_light")
	LightType_SpotLight      = LightType("spot_light")
	LightType_AmbientLight   = LightType("ambient_light")
)

type Vector3Wire struct {
	x float64 `json:"x" yaml:"x"`
	y float64 `json:"y" yaml:"y"`
	z float64 `json:"z" yaml:"z"`
}

type Material struct {
	Type  MaterialType `json:"type" yaml:"type"`
	Color Vector3Wire  `json:"color" yaml:"color"`
}

type MaterialSet struct {
	Ambient  Material `json:"ambient" yaml:"ambient"`
	Diffuse  Material `json:"diffuse" yaml:"diffuse"`
	Specular Material `json:"specular" yaml:"specular"`
}

type Primitive struct {
	Materials MaterialSet `json:"materials" yaml:"materials"`
}

type Sphere struct {
	Primitive `json:",inline" yaml:",inline"`

	Center Vector3Wire `json:"center" yaml:"center"`
	Radius float64     `json:"radius" yaml:"radius"`
}

type Triangle struct {
	Primitive `json:",inline" yaml:",inline"`

	VertexA Vector3Wire `json:"vertex_a" yaml:"vertex_a"`
	VertexB Vector3Wire `json:"vertex_b" yaml:"vertex_b"`
	VertexC Vector3Wire `json:"vertex_c" yaml:"vertex_c"`
}

type Plane struct {
	Primitive `json:",inline" yaml:",inline"`

	Normal Vector3Wire `json:"normal" yaml:"normal"`
	Depth  float64     `json:"depth" yaml:"depth"`
}

type Cylinder struct {
	Primitive `json:",inline" yaml:",inline"`

	Position Vector3Wire `json:"position" yaml:"position"`
	Axis     Vector3Wire `json:"axis" yaml:"axis"`
	Radius   float64     `json:"radius" yaml:"radius"`
	Length   float64     `json:"length" yaml:"length"`
}

type Light struct {
	Type      LightType   `json:"type" yaml:"type"`
	Center    Vector3Wire `json:"center" yaml:"center"`
	Direction Vector3Wire `json:"direction" yaml:"direction"`
	Color     Vector3Wire `json:"color" yaml:"color"`
}

type Camera struct {
	Center    Vector3Wire `json:"center" yaml:"center"`
	Direction Vector3Wire `json:"direction" yaml:"direction"`
}

type Image struct {
	Width  int `json:"width" yaml:"width"`
	Height int `json:"height" yaml:"height"`

	BackgroundColor Vector3Wire `json:"background" yaml:"background"`

	OutputFilePath string `json:"output_filepath" yaml:"output_filepath"`
}

type Config struct {
	Spheres   []Sphere   `json:"spheres" yaml:"spheres"`
	Triangles []Triangle `json:"triangles" yaml:"triangles"`
	Planes    []Plane    `json:"planes" yaml:"planes"`
	Cylinders []Cylinder `json:"cyliners" yaml:"cyliners"`

	Lights []Light `json:"lights" yaml:"lights"`

	MainCamera  Camera `json:"camera" yaml:"camera"`
	OutputImage Image  `json:"output_image" yaml:"output_image"`

	RecursionDepth int `json:"recursion_depth" yaml:"recursion_depth"`
}

func FromYaml(yamlFilePath string) (*Config, error) {
	data, err := os.ReadFile(yamlFilePath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
