package component

import (
	"github.com/Kahlis/rapina-go/structure/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type IComponent interface {
	Component()
}

var Empty Transform = Transform{
	Position: math.Vector3[float32]{X: 0.0, Y: 0.0, Z: 0.0},
	Rotation: math.Vector3[float32]{X: 0.0, Y: 0.0, Z: 0.0},
	Scale:    math.Vector3[float32]{X: 1.0, Y: 1.0, Z: 1.0},
}

type Transform struct {
	Position math.Vector3[float32]
	Rotation math.Vector3[float32]
	Scale    math.Vector3[float32]
}

func (Transform) Component() {}

func (t Transform) RlPosition() rl.Vector3 {
	return rl.NewVector3(
		t.Position.X,
		t.Position.Y,
		t.Position.Z,
	)
}
