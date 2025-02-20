package main

import (
	"github.com/MarcelArt/gonity/internal/entities"
	"github.com/MarcelArt/gonity/pkg/gameobject"
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

	gameObjects := start(entities.RegisterPlayerEntity)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		update(gameObjects)

		rl.EndDrawing()
	}
}

// func main() {
// 	rl.InitWindow(800, 450, "raylib [core] example - basic window")
// 	defer rl.CloseWindow()

// 	rl.SetTargetFPS(60)

// 	playerImage := rl.LoadImage("assets/character_squareGreen.png")
// 	playerTexture := rl.LoadTextureFromImage(playerImage)
// 	rl.UnloadImage(playerImage) // Unload image from RAM, it is now in VRAM
// 	for !rl.WindowShouldClose() {
// 		rl.BeginDrawing()

// 		rl.ClearBackground(rl.RayWhite)

// 		rl.DrawTexture(playerTexture, 600, 100, rl.White)

// 		rl.EndDrawing()
// 	}
// 	rl.UnloadTexture(playerTexture) // Unload texture when done
// }
