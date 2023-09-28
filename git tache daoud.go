package main

type Player struct {
	name              string
	class             string
	level             int
	lifemax           int
	life              int
	inventory         map[string]int
	money             int
	skill             []string
	equipment         Equipment
	inventoryUpgrades int
}

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Monstre struct {
	nom       string
	pvMax     int
	pvActuels int
	attaque   int
}
