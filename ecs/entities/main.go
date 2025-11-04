package entities

import (
	"github.com/Kahlis/rapina-go/ecs"
	"github.com/google/uuid"
)

var (
	Empty []Entity = []Entity{}
)

type Entity struct {
	Id         uuid.UUID
	Components map[int]ecs.Component
}

func (Entity) Entity() {}
