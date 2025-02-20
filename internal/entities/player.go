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

	playerMovement := &behaviours.PlayerMovement{
		Transform: transform,
		Speed:     20,
	}

	player.Components = append(player.Components, transform)
	player.Components = append(player.Components, sprite)
	player.Components = append(player.Components, playerMovement)

	return player
}
