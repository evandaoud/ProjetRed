package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

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

// //////////////////////////////////////tous les perso qui exisent
var p1 Player
var p2 Player
var p3 Player
var p4 Player

///////////////////////////////////FONCTION POUR LE MENU///////////////////////////////////////////////////

func Menu() {
	fmt.Println("Menu:")
	color.Blue("0. Choisir ton personnage")
	fmt.Println("-------------------------------------------")
	color.Blue("1. Afficher les informations du personnage")
	fmt.Println("-------------------------------------------")
	color.Blue("2. Afficher mon inventaire")
	fmt.Println("-------------------------------------------")
	color.Blue("3. Voir le marchand")
	fmt.Println("-------------------------------------------")
	color.Blue("4. Boire ma potion")
	fmt.Println("-------------------------------------------")
	color.Blue("5. Boire ma potion poison")
	fmt.Println("-------------------------------------------")
	color.Blue("6. Prendre ma boule de feu")
	fmt.Println("-------------------------------------------")
	color.Blue("7. Quitter")

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
		p1.AffInventory()
	case 3:
		fmt.Print("\033[H\033[2J")
		p1.Mymarchand()
	case 4:
		p1.TakePot()
	case 5:
		p1.PoisonPot()
	case 6:
		p1.SpellBook()
		Menu()

	case 7:
		fmt.Println("Au revoir !")
		os.Exit(0)
	default:
		fmt.Println("Option invalide.")
	}

}

///////////////////////////////////////////////////////FONCTION MAIN///////////////////////////////////////////////////

func main() {

	map1 := map[string]int{"Potions de vie": 3}
	p1.Init("Rayan", "Elfe", 1, 80, 40, map1, 100, "Coup de poing")

	p2.Init("Evan", "Demon", 1, 110, 50, map1, 100, "Coup de poing")

	p3.Init("Olivier", "Humain", 1, 100, 50, map1, 100, "Coup de poing")

	p4.Init("Jeremie", "Nain", 1, 120, 60, map1, 100, "Coup de poing")

}

/////////////////////////////////////////////////////CHOIX PERSO///////////////////////////////////////////////////////////

func ChoicePerso() {

	fmt.Println("Choissit ton personnage")
	fmt.Println("1-Humain")
	fmt.Println("2-Elfe")
	fmt.Println("3-Nain")

}

func (p *Player) Init(
	name string,
	class string,
	level int,
	lifemax int,
	life int,
	inventory map[string]int,
	money int,
	skill string,
) {

	p.name = name
	p.class = class
	p.level = level
	p.lifemax = lifemax
	p.life = life
	p.inventory = map[string]int{"potions": 1, "potion de poison": 1, "boule de feu": 1}
	p.money = money
	p.skill = []string{"Coup de poing"} // Ajoutez "Coup de poing" à la liste de compétences

	if skill != "" {
		p.skill = append(p.skill, skill) // Ajoutez le sort initial à la liste de compétences
	}
	p2.name = name
	p2.class = class
	p2.level = level
	p2.lifemax = lifemax
	p2.life = life
	p2.inventory = map[string]int{"potions": 1, "potion de poison": 1, "boule de feu": 1}
	p2.money = money
	p2.skill = []string{"Coup de poing"} // Ajoutez "Coup de poing" à la liste de compétences

	p3.name = name
	p3.class = class
	p3.level = level
	p3.lifemax = lifemax
	p3.life = life
	p3.inventory = map[string]int{"potions": 1, "potion de poison": 1, "boule de feu": 1}
	p3.money = money
	p3.skill = []string{"Coup de poing"} // Ajoutez "Coup de poing" à la liste de compétences

	p4.name = name
	p4.class = class
	p4.level = level
	p4.lifemax = lifemax
	p4.life = life
	p4.inventory = map[string]int{"potions": 1, "potion de poison": 1, "boule de feu": 1}
	p4.money = money
	p4.skill = []string{"Coup de poing"} // Ajoutez "Coup de poing" à la liste de compétences

}

// /////////////////////////////////////////////////Afficher information du personnage/////////////////////////////////////////////////////////
func (p Player) displayInfo() {
	fmt.Println("Informations du personnage :")
	color.Green("Nom : ", p.name)
	color.Green("Classe : ", p.class)
	color.Green("Niveau : ", p.level)
	color.Green("Vie maximale : ", p.lifemax)
	color.Green("Vie actuelle : ", p.life)
	color.Green("Inventaire : ", p.inventory)
	color.Green("Money : ", p.money)
	color.Green("Skill : ", p.skill)
	Menu()
}

///////////////////////////////////FONCTION BOULE DE FEU//////////////////////////////////////////////////////////////////////////////////////

func (p *Player) SpellBook() {
	newSpell := "Boule de feu"

	// Vérifiez si le sort n'est pas déjà dans la liste de compétences
	for _, spell := range p.skill {
		if spell == newSpell {
			color.Green("Vous avez déjà appris ce sort.%d")
			return
		}
	}

	/////////////////////////////////////////////////////// Ajoutez le sort à la liste de compétences///////////////////////////////////////////////
	p.skill = append(p.skill, newSpell)
	fmt.Printf("Vous avez appris le sort : %s\n", newSpell)
}

func (p *Player) Myboule() {
	fmt.Println("Tu as acheté une boule de feu")
	p.inventory["boule de feu"]++
	fmt.Print("\033[H\033[2J")

	// Appel de SpellBook pour ajouter le sort "Boule de feu" à la liste de compétences
	p.SpellBook()

	Menu()
}

///////////////////////////////////FONCTION BOIRE LE POISON///////////////////////////////////////////////////

func (p *Player) PoisonPot() {
	fmt.Println("Vous avez pris une potion poison")
	p1.Dead()
	fmt.Println("Vous avez :", p.life)
	for i := 0; i < 3; i++ {
		p.life -= 10
		fmt.Println("Vous perdez 10 PV")
		fmt.Println("Vous avez :", p.life, " / ", p.lifemax)
		time.Sleep(3 * time.Second)
		p.inventory["potion de poison"]--
		Menu()
	}

	color.Green("Vous avez maintenant : %d", p.life)
	time.Sleep(3 * time.Second)
	fmt.Print("\033[H\033[2J")
	Menu()
}

///////////////////////////////////FONCTION BOIRE LA POTION DE VIE///////////////////////////////////////////////////

func (p *Player) TakePot() {
	if p.inventory["potions"] > 0 {
		fmt.Println("Vous avez :", p.life)
		p.life += 50
		if p.life > p.lifemax {
			p.life = p.lifemax
		}
		fmt.Println("Vous avez bu une potion")
		color.Green("Vous avez maintenant : %d", p.life)
		time.Sleep(5000)
		p.inventory["potions"]--
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Println("Vous n'avez plus de potion")
	}
	Menu()

}

///////////////////////////////////FONCTION POUR L INVENTAIRE///////////////////////////////////////////////////

func (p *Player) AffInventory() {

	for item, count := range p.inventory {
		fmt.Println(item, " : ", count)
	}
	Menu()

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		os.Exit(1)
	}

	switch choice {
	case 1:
		p1.TakePot()
	case 2:
		p1.Mypoison()
	case 3:
		p1.Myboule()
	case 4:
		fmt.Println("Au revoir !")
		os.Exit(0)
	default:
		fmt.Println("Option invalide.")
	}

	fmt.Println("je te vends ca")
	Menu()
}

///////////////////////////////////FONCTION POUR LE MARCHAND///////////////////////////////////////////////////

func (p *Player) Mymarchand() {

	fmt.Println("Choisir un produit:")
	color.Green("Shop:")
	color.Green("1. Potions")
	color.Green("2. Potions de Poison")
	color.Green("3. Sort")
	fmt.Println()
	fmt.Println("Choisir un produit :")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		Menu()
	}

	switch choice {
	case 1:
		p1.Mypotion()
	case 2:
		p1.Mypoison()
	case 3:
		p1.Myboule()
	case 4:
		fmt.Println("Au revoir !")
		Menu()
	default:
		fmt.Println("Option invalide.")
	}

	fmt.Println("je te vends ca")
	Menu()

}

///////////////////////////////////FONCTION POUR ACHETER LA POTION///////////////////////////////////////////////////

func (p *Player) Mypotion() {

	if p.money >= 5 {
		fmt.Println("Tu as acheté une potion de vie à 5 elixir")
		p.Managemoney(-5)

		p.inventory["potions"]++
		fmt.Print("\033[H\033[2J")
		Menu()
	} else {
		println("Tu n'as pas assez d'oseilles")
	}
}

/////////////////////////////////// FONCTION POUR ACHETER LE POISON///////////////////////////////////////////////////

func (p *Player) Mypoison() {

	if p.money >= 5 {
		fmt.Println("Tu as acheté une potion de poison à 5 elixir")
		p.Managemoney(-5)
		fmt.Println("Tu as acheté une potion de de poison")
		p.inventory["potion de poison"]++
		fmt.Print("\033[H\033[2J")
	} else {
		fmt.Println("Tu n'as pas assez d'oseilles")
	}
	Menu()
}

// /////////////////////////////////////////////////////FONCTION POUR RESSUCITER///////////////////////////////////////////////////
func (p *Player) Dead() {
	if p.life <= 0 {
		fmt.Println("le joueur est mort")
		p.life += 50
		fmt.Println("Le joueur a été ressuscité avec 50% de points de vie maximum")
		Menu()
	}
}

// /////////////////////////////////////POUR ACHETER LA BOULE DE FEU QUE UNE FOIS ASKIP///////////////////////////////////////////////////
func (p *Player) Myboules() {
	if p.inventory["Sort"] > 0 {
		fmt.Println("Tu as acheté Un sort")
		p.inventory["boule de feu"]++
		fmt.Print("\033[H\033[2J")
	}
	Menu()
}

func (p *Player) Managemoney(nb int) {
	p.money += nb
}

//////////////////////////////////////////AMELIORATION DE LA CREATION DE PERSONNAGE////////////////////////////////////////////////////////////

func (p *Player) CharCreation() {

	fmt.Println("Choisir son le nom de ton personnage :")
	fmt.Println("le nom de ton personnage est :")
	fmt.Println("Choisit ton personnage :")
	fmt.Println("1) Humain")
	fmt.Println("2) Elfe")
	fmt.Println("3) Nain")

}

// //////////////////////////////////////////////////MONSTRE////////////////////////////////////////////////////////
type Monstre struct {
	nom      string
	pvmax    int
	pv       int
	pvAttack int
}

func (p *Monstre) InitGobbelin(nom string, pvmax int, pv int, pvAttack int) {

	p.nom = nom
	p.pvmax = pvmax
	p.pv = pv
	p.pvAttack = pvAttack

}

func (p *Monstre) TrainingFight() {
	var Gobelin Monstre
	Gobelin.InitGobbelin("Serge le Gobelin", 40, 40, 10)

}
