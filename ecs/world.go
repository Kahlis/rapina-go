package ecs

import (
	"fmt"
	"strconv"
	"time"
)

type World struct {
	FrameSnapshot [60]float64
	DeltaTime     time.Duration
	FrameTime     time.Time
	CurrentFrame  uint32
	FrameRate     uint32
	Entities      []Entity
	Systems       []System
}

func NewWorld(tps uint32, entities []Entity, systems []System) *World {
	w := &World{}
	return w.Create(tps, entities, systems)
}

func (w World) Create(frameRate uint32, entities []Entity, systems []System) *World {
	return &World{
		FrameSnapshot: [60]float64{},
		DeltaTime:     0.0,
		FrameTime:     time.Now(),
		CurrentFrame:  0,
		FrameRate:     frameRate,
		Entities:      entities,
		Systems:       systems,
	}
}

func (w *World) Run() {
	for i := range w.Systems {
		w.Systems[i].Run(w.CurrentFrame)
	}

	w.DeltaTime = time.Since(w.FrameTime)

	w.CurrentFrame++
	w.FrameTime = time.Now()
	w.FrameSnapshot[w.CurrentFrame%w.FrameRate] = 1 / w.DeltaTime.Seconds()
}

func (w *World) FPS(desiredFrames int) {
	if w.CurrentFrame%w.FrameRate == 0 {
		fps := 0.0
		for i := range w.FrameSnapshot {
			fps += w.FrameSnapshot[i]
		}
		fps /= float64(w.FrameRate)

		fpsTest := fmt.Sprintf("%.f", fps)

		if fpsTest != strconv.Itoa(desiredFrames) {
			frame, _ := strconv.Atoi(fpsTest)
			fmt.Printf("FrameGap: %d (%d)\n", frame, frame-desiredFrames)
		}
	}
}

func (w *World) WaitFrame() {
	frameDiff := (float64(w.FrameRate) - (1 / w.DeltaTime.Seconds())) * 4
	frameWait := time.Second / time.Duration(w.FrameRate)

	frameSkip := time.Duration(float64(frameWait) * (1 - frameDiff/100))

	if frameDiff > 0 {
		time.Sleep(frameSkip)
	} else {
		time.Sleep(frameWait)
	}
}

func (w *World) Exit() {}

func (w *World) AddEntity(e Entity) {
	w.Entities = append(w.Entities, e)
}

func (w *World) AddSystem(s System) {
	w.Systems = append(w.Systems, s)
}
