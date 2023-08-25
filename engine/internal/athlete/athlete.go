package athlete

type Stats struct {
	RunSpeed      int
	Vision        int
	TagDodgeSkill int
	Stealth       int
}

type Athlete struct {
	Name  string
	Stats Stats
}
