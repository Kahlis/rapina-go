package main

import (
	"github.com/Kahlis/rapina-go/ecs"
)

func main() {
	world := ecs.NewWorld(
		60,
		[]ecs.Entity{},
		[]ecs.System{},
	)

	for {
		world.Run()
	}
}
