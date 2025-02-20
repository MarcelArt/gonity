package gameobject

import (
	"log"

	"github.com/MarcelArt/gonity/pkg/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RigidBody2D struct {
	Velocity   rl.Vector2
	GameObject *GameObject

	transform *Transform2D
}

func (c *RigidBody2D) Start() {
	component := c.GameObject.GetComponent("Transform2D")

	transform, ok := (*component).(*Transform2D)
	if !ok {
		log.Println("Error: game object doesn't have Transform2D component")
	}

	c.transform = transform
}

func (c *RigidBody2D) Update() {
	deltaTime := rl.GetFrameTime()

	// Getting affected by gravity (if exists)
	c.Velocity.X += world.Gravity.X * deltaTime
	c.Velocity.Y += world.Gravity.Y * deltaTime

	// Affect transform
	c.transform.Position.X += c.Velocity.X * deltaTime
	c.transform.Position.Y += c.Velocity.Y * deltaTime
}

func (RigidBody2D) GetComponentName() string {
	return "RigidBody2D"
}
