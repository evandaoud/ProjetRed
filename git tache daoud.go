package red

type Player struct {
	name      string
	class     string
	level     int
	lifemax   int
	life      int
	inventory map[string]int
	money     int
	skill     []string
}

type Humain struct {
	name      string
	class     string
	level     int
	lifemax   int
	life      int
	inventory map[string]int
	money     int
	skill     []string
}

type Elfe struct {
	name      string
	class     string
	level     int
	lifemax   int
	life      int
	inventory map[string]int
	money     int
	skill     []string
}

type Nain struct {
	name      string
	class     string
	level     int
	lifemax   int
	life      int
	inventory map[string]int
	money     int
	skill     []string
}
