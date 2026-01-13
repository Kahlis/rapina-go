package position

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Zero Position = Position{X: 0.0, Y: 0.0, Z: 0.0}

type Position struct {
	X float32
	Y float32
	Z float32
}

func (p Position) RlPosition() rl.Vector3 {
	return rl.NewVector3(p.X, p.Y, p.Z)
}

func (Position) Component() {}
