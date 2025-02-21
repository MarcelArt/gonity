package main

import (
	"github.com/MarcelArt/gonity/internal/entities"
	"github.com/MarcelArt/gonity/pkg/gameobject"
	"github.com/MarcelArt/gonity/pkg/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func start(registerFuncs ...func() *gameobject.GameObject) []*gameobject.GameObject {
	gameObjects := make([]*gameobject.GameObject, 0)
	for _, registerFunc := range registerFuncs {
		gameObject := registerFunc()
		gameObjects = append(gameObjects, gameObject)

		for _, component := range gameObject.Components {
			component.Start()
		}
	}
	return gameObjects
}

func update(gameObjects []*gameobject.GameObject) {
	for _, gameObject := range gameObjects {
		for _, component := range gameObject.Components {
			component.Update()
		}
	}
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	world.SetupGravity(0, 980)

	gameObjects := start(entities.RegisterPlayerEntity, entities.RegisterMapEntity)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawFPS(0, 0)

		rl.ClearBackground(rl.Green)

		update(gameObjects)

		rl.EndDrawing()
	}
}
