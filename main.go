package main

import (
	"fmt"
	"os"
)

type Player struct {
	name      string
	class     string
	level     int
	lifemax   int
	life      int
	inventory map[string]int
	money     int
}

type Inventory struct {
	drunkpotion string
}

var p1 Player

func main() {
	fmt.Println("Menu:")
	fmt.Println("1. Afficher les informations du personnage")
	fmt.Println("2. Afficher a ma potion")
	fmt.Println("3. QUittez")

	var choice int
	fmt.Print("Choisissez une option : ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		os.Exit(1)
	}

	switch choice {
	case 1:
		p1.displayInfo()
	case 2:
		p1.TakePot()
	case 3:
		fmt.Println("Au revoir !")
		os.Exit(0)
	default:
		fmt.Println("Option invalide.")
	}
}
func (p *Player) Init(
	name string,
	class string,
	level int,
	lifemax int,
	life int,
	inventory map[string]int,
	money int,
) {
	p.name = name
	p.class = class
	p.level = level
	p.lifemax = lifemax
	p.life = life
	p.inventory = inventory
	p.money = money

}
func (p Player) displayInfo() {
	fmt.Println("Informations du personnage :")
	fmt.Println("Nom :", p.name)
	fmt.Println("Classe :", p.class)
	fmt.Println("Niveau :", p.level)
	fmt.Println("Vie maximale :", p.lifemax)
	fmt.Println("Vie actuelle :", p.life)
	fmt.Println("Inventaire :", p.inventory)
	fmt.Println("money :", p.money)
}

func (p *Player) TakePot() {
	fmt.Println("Vous avez : ", p.life)
	p.life = p.life + 50
	fmt.Println("Vous avez pris une potion")
	fmt.Println("Vous avez maintenant :", p.life)
}

func (p *Player) AffInventory() {
	//créer un tableau vide
	//var tableau []string

	//boucler dans ta map
	//envoyer ce que tu trouve dans le tableau
	// affiche ton tableau
	// si ta map est vide tu print que l'inventaire est vide
	//afficher les choix (prendre la potion ou non)
	//refaire un scanln
}
