package stance

type Stance string

func (t Stance) String() string {
	return string(t)
}

func FromString(v string) Stance {
	if e, ok := valueMap[v]; ok {
		return e
	}
	return NotSet
}

const (
	NotSet = Stance("NOT_SET")
	Run    = Stance("RUN")
	Scout  = Stance("SCOUT")
	Crawl  = Stance("CRAWL")
	Walk   = Stance("WALK")
)

var valueMap = map[string]Stance{
	"NOT_SET": Stance("NOT_SET"),
	"RUN":     Stance("RUN"),
	"SCOUT":   Stance("SCOUT"),
	"CRAWL":   Stance("CRAWL"),
	"WALK":    Stance("WALK"),
}
