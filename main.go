package main

import (
	"github.com/Kahlis/rapina-go/ecs"
	"github.com/Kahlis/rapina-go/ecs/component"
	"github.com/Kahlis/rapina-go/ecs/entity"
	"github.com/Kahlis/rapina-go/ecs/system"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

/*
// World Example
func main() {
	world := ecs.NewWorld(
		60,
		[]ecs.Entity{},
		[]ecs.System{},
	)

	for {
		world.Run()
		world.WaitFrame()
	}
}
*/

func main() {
	cubeTransform := &component.Empty
	cubeEntity := entity.Entity{
		Id:         uuid.New(),
		Components: map[int]component.IComponent{0: cubeTransform},
	}

	moverSystem := system.Mover{}

	world := ecs.NewWorld(
		60,
		[]entity.Entity{cubeEntity},
		[]system.ISystem{moverSystem},
	)

	rl.InitWindow(800, 450, "Rapina Engine [core] example - 3d camera free")

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(10.0, 10.0, 10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	// cubePosition := rl.NewVector3(cubeTransform.Position.X, 0.0, 0.0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree) // Update camera with free camera mode

		if rl.IsKeyDown(rl.KeyZ) {
			camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
		}

		world.Run()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(cubeTransform.RlPosition(), 2.0, 2.0, 2.0, rl.Red)
		rl.DrawCubeWires(cubeTransform.RlPosition(), 2.0, 2.0, 2.0, rl.Maroon)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 320, 133, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 320, 133, rl.Blue)

		rl.DrawText("Free camera default controls:", 20, 20, 10, rl.Black)
		rl.DrawText("- Mouse Wheel to Zoom in-out", 40, 40, 10, rl.DarkGray)
		rl.DrawText("- Mouse Wheel Pressed to Pan", 40, 60, 10, rl.DarkGray)
		rl.DrawText("- Z to zoom to (0, 0, 0)", 40, 120, 10, rl.DarkGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
