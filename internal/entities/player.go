package entities

import (
	"github.com/MarcelArt/gonity/internal/behaviours"
	"github.com/MarcelArt/gonity/pkg/gameobject"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func RegisterPlayerEntity() *gameobject.GameObject {
	player := &gameobject.GameObject{
		Name: "Player",
	}

	transform := &gameobject.Transform2D{
		Position: rl.Vector2{X: 600, Y: 100},
		Scale:    rl.Vector2Zero(),
	}

	sprite := &gameobject.Sprite2D{
		ImagePath:  "assets/character_squareGreen.png",
		Tint:       rl.White,
		GameObject: player,
	}

	rigidBody := &gameobject.RigidBody2D{
		Velocity:   rl.Vector2Zero(),
		GameObject: player,
	}

	playerMovement := &behaviours.PlayerMovement{
		Transform:    transform,
		RigidBody:    rigidBody,
		Speed:        20,
		JumpStrength: 100,
	}

	player.Components = append(player.Components, transform)
	player.Components = append(player.Components, sprite)
	player.Components = append(player.Components, playerMovement)
	player.Components = append(player.Components, rigidBody)

	return player
}
