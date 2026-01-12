package components

import "github.com/Kahlis/rapina-go/structure/math"

type Transform struct {
	Position math.Vector2[int]
	Rotation float32
	Scale    math.Vector2[float32]
}

func (Transform) Component() {}
