package system

import (
	"fmt"
	"math"

	"github.com/Kahlis/rapina-go/ecs/component"
	"github.com/Kahlis/rapina-go/ecs/entity"
)

type Mover struct {
	System
}

func (c Mover) Init() {
	fmt.Println("Cube Mover initialized!")
}

func (c Mover) Run(frame uint32, entities []entity.Entity) {
	if len(entities) == 0 {
		return
	}

	if len(entities[0].Components) == 0 {
		return
	}

	cp := entities[0].Components[0]
	if transform, ok := cp.(*component.Transform); ok {
		time := float64(frame) / (60.0 / 4) // 4 cycles per second at 60 fps
		transform.Position.X = float32(math.Cos(time) * 2)
	}
}

func (c Mover) Exit() {
	fmt.Println("Cube Translator destroyed!")
}
