package athlete

// What about captains?  Do they get the same stats

type Stats struct {
	RunSpeed      int
	Vision        int
	TagDodgeSkill int
	Stealth       int
	Climb         int
	Swim          int
}

type Athlete struct {
	Name  string
	Stats Stats
}
