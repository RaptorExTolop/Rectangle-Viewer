package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Rectangle struct {
	Pos, Size rl.Vector2
	Colour rl.Color
	Angle float32

	_Rectangle rl.Rectangle
}

func (r *Rectangle) Init(Pos, Size rl.Vector2, Colour rl.Color, Angle float32) {
	r.Pos = Pos
	r.Size = Size
	r.Colour = Colour
	r.Angle = Angle

	r._Rectangle = rl.NewRectangle(0, 0, r.Size.X, r.Size.Y)
}

func (r *Rectangle) Draw() {
	// rl.DrawRectanglePro(r._Rectangle, rl.Vector2Zero(), 0, r.Colour)
	rl.DrawRectanglePro(r._Rectangle, r.Pos, r.Angle, r.Colour)
}
