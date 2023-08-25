package gamemap

import (
	"testing"

	"github.com/menwald/ctf/engine/internal/enum/terraintype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMap(t *testing.T) {
	t.Run("TestNewMap_SmallForestOk", func(t *testing.T) {
		width := 20
		height := 10
		gm, err := NewGameMap(width, height, func(x int, y int) Terrain {
			return Terrain{Type: terraintype.Forest}
		})
		require.NoError(t, err)

		for w := 0; w < width; w++ {
			for h := 0; h < height; h++ {
				ter, err := gm.Terrain(Coordinates{X: w, Y: h})
				require.NoError(t, err)
				assert.Equal(t, ter.Type, terraintype.Forest)
			}
		}
	})

	t.Run("TestNewMap_RowsOk", func(t *testing.T) {
		width := 20
		height := 10
		gm, err := NewGameMap(width, height, func(x int, y int) Terrain {
			switch y {
			case 0:
				return Terrain{Type: terraintype.NotSet}
			case 1:
				return Terrain{Type: terraintype.PavedRoad}
			case 2:
				return Terrain{Type: terraintype.DirtPath}
			case 3:
				return Terrain{Type: terraintype.Grass}
			case 4:
				return Terrain{Type: terraintype.Forest}
			default:
				return Terrain{Type: terraintype.Water}
			}
		})
		require.NoError(t, err)

		for w := 0; w < width; w++ {
			ter, err := gm.Terrain(Coordinates{X: w, Y: 0})
			require.NoError(t, err)
			assert.Equal(t, ter.Type, terraintype.NotSet)

			ter, err = gm.Terrain(Coordinates{X: w, Y: 1})
			require.NoError(t, err)
			assert.Equal(t, ter.Type, terraintype.PavedRoad)

			ter, err = gm.Terrain(Coordinates{X: w, Y: 2})
			require.NoError(t, err)
			assert.Equal(t, ter.Type, terraintype.DirtPath)

			ter, err = gm.Terrain(Coordinates{X: w, Y: 3})
			require.NoError(t, err)
			assert.Equal(t, ter.Type, terraintype.Grass)

			ter, err = gm.Terrain(Coordinates{X: w, Y: 4})
			require.NoError(t, err)
			assert.Equal(t, ter.Type, terraintype.Forest)
			for h := 5; h < height; h++ {
				ter, err = gm.Terrain(Coordinates{X: w, Y: h})
				require.NoError(t, err)
				assert.Equal(t, ter.Type, terraintype.Water)
			}
		}
	})
}
