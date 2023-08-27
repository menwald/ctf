package gamemap

import (
	"fmt"

	"github.com/menwald/ctf/engine/internal/enum/terraintype"
)

type Coordinates struct {
	X int
	Y int
}

type GameMap struct {
	terrain [][]Terrain
}

func NewGameMap(width int, height int, tfun func(int, int) Terrain) (GameMap, error) {
	m := GameMap{terrain: make([][]Terrain, width)}
	for i := range m.terrain {
		m.terrain[i] = make([]Terrain, height)
		for j := range m.terrain[i] {
			m.terrain[i][j] = tfun(i, j)
		}
	}
	return m, nil
}

func (g GameMap) Terrain(c Coordinates) (Terrain, error) {
	if !g.areValid(c) {
		return Terrain{Type: terraintype.NotSet, Obstacles: nil}, fmt.Errorf("invalid coordinates")
	}
	return g.terrain[c.X][c.Y], nil
}

func (g GameMap) areValid(c Coordinates) bool {
	return len(g.terrain) > c.X && len(g.terrain[c.X]) > c.Y
}

func Generate(width uint, height uint, maxTeams uint) (GameMap, error) {
	m := GameMap{terrain: make([][]Terrain, width)}
	for i := range m.terrain {
		m.terrain[i] = make([]Terrain, height)
		for j := range m.terrain[i] {
			m.terrain[i][j] = Terrain{
				Type:            terraintype.Forest,
				Obstacles:       nil,
				RunResistance:   3,
				ClimbResistance: 0,
				SwimResistance:  0,
			}
		}
	}

	return m, nil
}
