package entity

import (
	"github.com/Kahlis/rapina-go/ecs/component"
	"github.com/google/uuid"
)

var (
	Empty []Entity = []Entity{}
)

type Entity struct {
	Id         uuid.UUID
	Components map[int]component.IComponent
}

func (Entity) Entity() {}
