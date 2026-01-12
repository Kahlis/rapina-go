package system

import (
	"github.com/Kahlis/rapina-go/ecs/entity"
	"github.com/google/uuid"
)

var (
	Empty []System = []System{}
)

type ISystem interface {
	Init()
	Run(uint32, []entity.Entity)
	Exit()
}

type System struct {
	Id uuid.UUID
}
