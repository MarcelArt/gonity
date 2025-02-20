package gameobject

import (
	"image/color"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite2D struct {
	ImagePath  string
	Tint       color.RGBA
	GameObject *GameObject
	image      rl.Image
	texture    rl.Texture2D
}

func (c *Sprite2D) Start() {
	c.image = *rl.LoadImage(c.ImagePath)

	c.texture = rl.LoadTextureFromImage(&c.image)
	log.Println("Texture loaded")

	rl.UnloadImage(&c.image)
}

func (Sprite2D) GetComponentName() string {
	return "Sprite2D"
}

func (c *Sprite2D) Update() {
	component := c.GameObject.GetComponent("Transform2D")

	transform, ok := (*component).(*Transform2D)
	if !ok {
		log.Println("Error: game object doesn't have Transform2D component")
		return
	}

	rl.DrawTextureV(c.texture, transform.Position, c.Tint)
}
