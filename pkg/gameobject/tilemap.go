package gameobject

import rl "github.com/gen2brain/raylib-go/raylib"

type Tile struct {
	X         int
	Y         int
	ImagePath string
	texture   rl.Texture2D
}

type Tilemap struct {
	TileWidth  float32
	TileHeight float32
	Tiles      []*Tile
}

func (c *Tilemap) Start() {
	for _, tile := range c.Tiles {
		tile.texture = rl.LoadTexture(tile.ImagePath)
	}
}

func (c *Tilemap) Update() {
	for _, tile := range c.Tiles {
		var x float32 = c.TileWidth * float32(tile.X)
		var y float32 = c.TileHeight * float32(tile.Y)
		rl.DrawTextureV(tile.texture, rl.NewVector2(x, y), rl.White)
	}
}

func (c *Tilemap) GetComponentName() string {
	return "Tilemap"
}
