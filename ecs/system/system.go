package system

import (
	"time"

	"github.com/google/uuid"
)

var (
	Empty []System = []System{}
)

type ISystem interface {
	Init()
	Run(uint32, time.Duration)
	Exit()
}

type System struct {
	Id uuid.UUID
}
