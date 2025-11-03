package systems

import "github.com/Kahlis/rapina-go/ecs"

type World struct {
	Entities []ecs.Entity
	Systems  []ecs.System
}
