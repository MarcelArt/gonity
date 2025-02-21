package entities

import "github.com/MarcelArt/gonity/pkg/gameobject"

func RegisterMapEntity() *gameobject.GameObject {
	levelMap := &gameobject.GameObject{
		Name: "Map",
	}

	tileMap := &gameobject.Tilemap{
		TileWidth:  61,
		TileHeight: 61,
		Tiles: []*gameobject.Tile{
			{
				X:         0,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         1,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         2,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         3,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         4,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         5,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         6,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         7,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         8,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         9,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         10,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         11,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
			{
				X:         12,
				Y:         6,
				ImagePath: "assets/tile_grass.png",
			},
		},
	}

	levelMap.Components = append(levelMap.Components, tileMap)

	return levelMap
}
