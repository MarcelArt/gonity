package behaviours

import (
	"github.com/MarcelArt/gonity/pkg/gameobject"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerMovement struct {
	Transform    *gameobject.Transform2D
	RigidBody    *gameobject.RigidBody2D
	Speed        float32
	JumpStrength float32
}

func (c *PlayerMovement) Start() {

}

func (c *PlayerMovement) Update() {
	if rl.IsKeyDown(rl.KeyLeft) {
		c.Transform.Position.X -= c.Speed
	}
	if rl.IsKeyDown(rl.KeyRight) {
		c.Transform.Position.X += c.Speed
	}
	if rl.IsKeyDown(rl.KeyZ) {
		c.RigidBody.Velocity.Y = -c.JumpStrength
	}
}

func (PlayerMovement) GetComponentName() string {
	return "PlayerMovement"
}
