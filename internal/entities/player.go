package entities

import (
	"github.com/MarcelArt/gonity/pkg/gameobject"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func RegisterPlayerEntity() *gameobject.GameObject {
	player := &gameobject.GameObject{
		Name: "Player",
	}

	player.Components = append(player.Components, &gameobject.Transform2D{
		Position: rl.Vector2{X: 600, Y: 100},
		Scale:    rl.Vector2Zero(),
	})
	player.Components = append(player.Components, &gameobject.Sprite2D{
		ImagePath:  "assets/character_squareGreen.png",
		Tint:       rl.White,
		GameObject: player,
	})

	return player
}
