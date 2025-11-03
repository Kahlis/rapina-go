package entities

import (
	"github.com/Kahlis/rapina-go/ecs"
	"github.com/google/uuid"
)

type Entity struct {
	Id         uuid.UUID
	Components []ecs.Component
}

func (Entity) Entity() {}
