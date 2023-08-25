package terraintype

type TerrainType string

func (t TerrainType) String() string {
	return string(t)
}

func FromString(v string) TerrainType {
	if e, ok := valueMap[v]; ok {
		return e
	}
	return NotSet
}

const (
	NotSet    = TerrainType("NOT_SET")
	PavedRoad = TerrainType("PAVED_ROAD")
	DirtPath  = TerrainType("DIRT_PATH")
	Grass     = TerrainType("GRASS")
	Forest    = TerrainType("FOREST")
	Water     = TerrainType("WATER")
)

var valueMap = map[string]TerrainType{
	"NOT_SET":    TerrainType("NOT_SET"),
	"PAVED_ROAD": TerrainType("PAVED_ROAD"),
	"DIRT_PATH":  TerrainType("DIRT_PATH"),
	"GRASS":      TerrainType("GRASS"),
	"FOREST":     TerrainType("FOREST"),
	"WATER":      TerrainType("WATER"),
}
