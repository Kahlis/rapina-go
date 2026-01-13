package math

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

type Vector2[T Number] struct {
	X, Y T
}

type Vector3[T Number] struct {
	X, Y, Z T
}
