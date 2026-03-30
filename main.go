package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ( /* Special Global Variables */
	// super awesome and cool window and game settings
	WINDOW_WIDTH, WINDOW_HEIGHT int32 = 1280, 720
	WINDOW_TITLE string = "rectangle position getter"
	TARGET_FPS int32 = 60
	RUNNING bool = true

	BKG_COLOUR rl.Color = rl.SkyBlue
)

var ( /* not so special global variables*/
	rec1 Rectangle
)

func init() {
  rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, WINDOW_TITLE)
  rl.SetTargetFPS(TARGET_FPS)

  rec1.Init(rl.NewVector2(100, 100), rl.NewVector2(300, 150), rl.Maroon, 0)

  fmt.Println(rec1)
}

func main() {
  defer rl.CloseWindow()

  for RUNNING {
	  update()
	  draw()
  }
}

func update() {
	RUNNING = !rl.WindowShouldClose()

	if (rl.IsKeyDown(rl.KeyE)) {
		rec1.Angle+=16 * rl.GetFrameTime()
	}
}

func draw() {
	rl.BeginDrawing()

    rl.ClearBackground(BKG_COLOUR)
	rl.DrawFPS(0, 0)

	rec1.Draw()

    rl.EndDrawing()
}
