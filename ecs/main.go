package ecs

type Component interface {
	Component()
}

type Entity interface {
	Entity()
}

type System interface {
	Init()
	Run(uint32)
	Exit()
}
