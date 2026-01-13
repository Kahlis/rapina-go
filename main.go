package main

import (
	"fmt"

	"github.com/Kahlis/rapina-go/ecs"
	"github.com/Kahlis/rapina-go/ecs/component"
	"github.com/Kahlis/rapina-go/ecs/component/position"
	"github.com/Kahlis/rapina-go/ecs/entity"
	"github.com/Kahlis/rapina-go/ecs/system"
	"github.com/Kahlis/rapina-go/ecs/system/core"
	"github.com/Kahlis/rapina-go/structure/input"
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
	cubePosition := &position.Zero
	keys := make(map[int32]input.KeyInterest)
	input := input.Keys{List: keys}

	input.BulkSet([]int32{
		rl.KeyLeft,
		rl.KeyRight,
		rl.KeyUp,
		rl.KeyDown,
	})

	cubeEntity := entity.Entity{
		Id:         uuid.New(),
		Components: map[int]component.IComponent{0: cubePosition},
	}

	moverSystem := core.Motor{}

	world := ecs.NewWorld(
		60,
		[]entity.Entity{cubeEntity},
		[]system.ISystem{moverSystem},
		input,
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
	toggle := true
	disabled := ""

	for !rl.WindowShouldClose() {
		if rl.IsMouseButtonDown(rl.MouseRightButton) {
			rl.UpdateCamera(&camera, rl.CameraFree)
			rl.DisableCursor()
		} else if rl.IsCursorHidden() {
			rl.EnableCursor()
		}

		if rl.IsKeyDown(rl.KeyZ) {
			camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			if toggle {
				input.Delete(rl.KeyLeft)
				toggle = false
				disabled = " (Disabled)"
			} else {
				input.Set(rl.KeyLeft)
				toggle = true
				disabled = ""
			}
		}

		world.Run()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(cubePosition.RlPosition(), 2.0, 2.0, 2.0, rl.Red)
		rl.DrawCubeWires(cubePosition.RlPosition(), 2.0, 2.0, 2.0, rl.Maroon)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()

		rl.DrawRectangle(10, 10, 320, 133, rl.Fade(rl.SkyBlue, 0.5))
		rl.DrawRectangleLines(10, 10, 320, 133, rl.Blue)

		rl.DrawText("Input Test:", 20, 20, 10, rl.Black)
		rl.DrawText(fmt.Sprintf("LeftArrow%s: %t", disabled, input.GetKey(rl.KeyLeft).Down), 40, 40, 10, rl.DarkGray)
		rl.DrawText(fmt.Sprintf("RightArrow: %t", input.GetKey(rl.KeyRight).Down), 40, 60, 10, rl.DarkGray)
		rl.DrawText(fmt.Sprintf("UpArrow: %t", input.GetKey(rl.KeyUp).Down), 40, 80, 10, rl.DarkGray)
		rl.DrawText(fmt.Sprintf("DownArrow: %t", input.GetKey(rl.KeyDown).Down), 40, 100, 10, rl.DarkGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
