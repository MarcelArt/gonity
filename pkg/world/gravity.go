package world

import rl "github.com/gen2brain/raylib-go/raylib"

var Gravity rl.Vector2

func SetupGravity(x float32, y float32) {
	Gravity = rl.Vector2{
		X: x,
		Y: y,
	}
}
