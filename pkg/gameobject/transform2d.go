package gameobject

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Transform2D struct {
	Position rl.Vector2
	Scale    rl.Vector2
}

func (c *Transform2D) Start() {
	// log.Printf("Starting (X, Y): (%f, %f)", c.Position.X, c.Position.Y)
}

func (c *Transform2D) Update() {
	// c.Position.X += 10
	// c.Position.Y += 1
	// log.Printf("Current (X, Y): (%f, %f)", c.Position.X, c.Position.Y)
}

func (Transform2D) GetComponentName() string {
	return "Transform2D"
}
