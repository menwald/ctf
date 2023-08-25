package gamemap

import (
	"github.com/menwald/ctf/engine/internal/enum/terraintype"
)

type Obstacle struct {
	// fences, etc?
}

type Terrain struct {
	Type      terraintype.TerrainType
	Obstacles []Obstacle
}
