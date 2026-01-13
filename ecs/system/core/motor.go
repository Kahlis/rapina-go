package core

import (
	"fmt"
	"math"

	"github.com/Kahlis/rapina-go/ecs/component/position"
	"github.com/Kahlis/rapina-go/ecs/entity"
)

type Motor struct{}

func (c Motor) Init() {
	fmt.Println("Cube Mover initialized!")
}

func (c Motor) Run(frame uint32, entities []entity.Entity) {
	if len(entities) == 0 {
		return
	}

	if len(entities[0].Components) == 0 {
		return
	}

	cp := entities[0].Components[0]
	if position, ok := cp.(*position.Position); ok {
		time := float64(frame) / (60.0 / 4) // 4 cycles per second at 60 fps
		position.X = float32(math.Cos(time) * 2)
	}
}

func (c Motor) Exit() {
	fmt.Println("Cube Translator destroyed!")
}
