package core

import (
	"fmt"
	"time"

	"github.com/Kahlis/rapina-go/ecs/component/position"
	"github.com/Kahlis/rapina-go/ecs/entity"
	"github.com/Kahlis/rapina-go/structure/input"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

type Motor struct {
	Input    input.Keys
	Entities map[uuid.UUID]entity.Entity
	Target   uuid.UUID
}

func (c Motor) Init() {
	fmt.Println("Cube Mover initialized!")
}

func (c Motor) Run(frame uint32, delta time.Duration) {
	cp := c.Entities[c.Target].Components[0]
	if position, ok := cp.(*position.Position); ok {
		var horizontal float32 = 0.0
		var vertical float32 = 0.0

		if c.Input.GetKey(rl.KeyRight).Down {
			horizontal += 1.0
		}

		if c.Input.GetKey(rl.KeyLeft).Down {
			horizontal -= 1.0
		}

		if c.Input.GetKey(rl.KeyUp).Down {
			vertical += 1.0
		}

		if c.Input.GetKey(rl.KeyDown).Down {
			vertical -= 1.0
		}

		movement := rl.Vector3{X: horizontal, Y: 0.0, Z: -vertical}
		newPosition := rl.Vector3Add(
			position.RlPosition(),
			rl.Vector3Scale(
				movement,
				float32(delta.Seconds())*10,
			),
		)

		position.X, position.Z = newPosition.X, newPosition.Z
	}
}

func (c Motor) Exit() {
	fmt.Println("Cube Translator destroyed!")
}
