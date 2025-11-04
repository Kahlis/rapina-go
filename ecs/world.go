package ecs

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type World struct {
	DeltaTime    time.Duration
	FrameTime    time.Time
	CurrentFrame uint32
	FrameRate    uint8
	Entities     []Entity
	Systems      []System
}

func NewWorld(tps uint8, entities []Entity, systems []System) *World {
	w := &World{}
	return w.Create(tps, entities, systems)
}

func (w World) Create(frameRate uint8, entities []Entity, systems []System) *World {
	return &World{
		DeltaTime:    0,
		FrameTime:    time.Now(),
		CurrentFrame: 0,
		FrameRate:    frameRate,
		Entities:     entities,
		Systems:      systems,
	}
}

func (w *World) Run() {
	for i := range w.Systems {
		w.Systems[i].Run(w.CurrentFrame)
	}

	w.DeltaTime = time.Since(w.FrameTime)

	w.CurrentFrame++
	w.FrameTime = time.Now()

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Printf("FrameRate is %f\n", 1/w.DeltaTime.Seconds())
}

func (w *World) Exit() {}

func (w *World) AddEntity(e Entity) {
	w.Entities = append(w.Entities, e)
}

func (w *World) AddSystem(s System) {
	w.Systems = append(w.Systems, s)
}
