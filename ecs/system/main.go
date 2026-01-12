package systems

import "github.com/google/uuid"

var (
	Empty []System = []System{}
)

type System struct {
	Id uuid.UUID
}
