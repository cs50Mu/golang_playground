package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SPEED         = 300
	SCREEN_FACTOR = 80
	SCREEN_WIDTH  = 16 * SCREEN_FACTOR
	SCREEN_HEIGHT = 9 * SCREEN_FACTOR
	BRANCH_COUNT  = 5
	BRANCH_LENGTH = SCREEN_FACTOR * 2
	BRANCH_ANGLE  = 2 * math.Pi / BRANCH_COUNT
	BRANCH_THICK  = 10.0
)

func drawBranch(center rl.Vector2, branchLen int,
	thick float32, hue float32, depth int) {
	if depth <= 0 {
		return
	}

	for i := 0; i < BRANCH_COUNT; i++ {
		branch := rl.Vector2{
			X: center.X + float32(branchLen)*float32(math.Cos(BRANCH_ANGLE*float64(i))),
			Y: center.Y + float32(branchLen)*float32(math.Sin(BRANCH_ANGLE*float64(i))),
		}
		color := rl.ColorFromHSV(hue, 1.0, 1.0)
		rl.DrawLineEx(
			center,
			branch,
			thick,
			color,
		)
		drawBranch(branch, branchLen/2, thick*0.5, hue+70.0, depth-1)
	}
}

func snowflake() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "snowflake")
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(0x18, 0x18, 0x18, 0xff))
		center := rl.Vector2{
			X: float32(rl.GetScreenWidth()) * 0.5,
			Y: float32(rl.GetScreenHeight()) * 0.5,
		}
		drawBranch(center, BRANCH_LENGTH, BRANCH_THICK, 0, 4)
		rl.EndDrawing()
	}
}

func main() {
	snowflake()
}

func demo() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "hello, world")
	defer rl.CloseWindow()

	// rl.SetTargetFPS(60)

	size := rl.Vector2{X: 100, Y: 100}
	position := rl.Vector2{X: 0, Y: 0}
	velocity := rl.Vector2{X: SPEED, Y: SPEED}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.NewColor(0x18, 0x18, 0x18, 0xff))

		position.X += velocity.X * rl.GetFrameTime()
		position.Y += velocity.Y * rl.GetFrameTime()

		rl.DrawRectangleV(
			position,
			size,
			rl.Red)

		rl.EndDrawing()
	}
}
