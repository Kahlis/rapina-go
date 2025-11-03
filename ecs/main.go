package ecs

type Component interface {
	Component()
}

type Entity interface {
	Entity()
}

type System interface {
	Start()
	Run(uint32)
	Exit()
}
