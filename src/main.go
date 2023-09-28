package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

// ////////////////////////////////////////////SCROLL LE TEXTE//////////////////////////////////////
func SpeedMsg(message string, speed int, colorName string) {
	defaultColor := color.New(color.FgWhite)
	Green := color.New(color.FgGreen)
	Red := color.New(color.FgRed)
	Blue := color.New(color.FgBlue)
	Cyan := color.New(color.FgCyan)
	Yellow := color.New(color.FgYellow)

	var selectedColor *color.Color

	switch colorName {
	case "green":
		selectedColor = Green
	case "red":
		selectedColor = Red
	case "blue":
		selectedColor = Blue
	case "cyan":
		selectedColor = Cyan
	case "yellow":
		selectedColor = Yellow
	default:
		selectedColor = defaultColor
	}

	for _, char := range message {
		selectedColor.Print(string(char))
		time.Sleep(time.Duration(10) * time.Millisecond)
	}
}

// ////////////////////////////////////////////////////////////Joueur/////////////////////////
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

// /////////////////////////////////////////////////////////////MONSTRE////////////////////////////////////////////////////////////////////////
type Monstre struct {
	Name      string
	MaxHealth int
	Health    int
	Attack    int
}

// /////////////////////////////////////////////////////////////////////////Fonction Gobelin///////////////////////////////////////////////////
type Character struct {
	Name      string
	Health    int
	MaxHealth int
	Attack    int
}

func GoblinPattern(player *Character, goblin *Monstre) {
	rand.Seed(time.Now().UnixNano())

	// Initialisation des variables
	playerHealth := player.Health
	goblinHealth := goblin.Health
	turnCounter := 0

	// Boucle de combat
	for playerHealth > 0 && goblinHealth > 0 {
		turnCounter++
		if turnCounter%3 == 0 {
			// Tous les 3 tours, le Gobelin inflige 200% de son attaque en dégâts
			damage := 2 * goblin.Attack
			playerHealth -= damage
		} else {
			// Sinon, le Gobelin inflige 100% de son attaque en dégâts
			playerHealth -= goblin.Attack
		}

		// Affichage des dégâts infligés
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, player.Name, goblin.Attack)

		// Affichage des points de vie actuels du joueur
		fmt.Printf("Points de vie actuels de %s: %d/%d\n", player.Name, playerHealth, player.MaxHealth)
	}

	// Vérification du résultat du combat
	if playerHealth <= 0 {
		fmt.Printf("%s a été vaincu par le %s.\n", player.Name, goblin.Name)
	} else {
		fmt.Printf("%s a vaincu le %s.\n", player.Name, goblin.Name)
	}
}

///////////////////////////////////////////////////////////////////////////////////RACONTER HISTOIRE//////////////////////////////////////

func RaconterHistoire() {
	ClearConsole()

	SpeedMsg("\t   ____             U _____ u_   _  __     __ U _____ u_   _       _   _U _____ u    \n\tU | __\")u    ___    \\| ___\"|/ \\ |\"| \\ \\   /\"/u\\| ___\"|/ \\ |\"|   U |\"|u| \\| ___\"|/    \n\t \\|  _ \\/   |_\"_|    |  _|\"<|  \\| |> \\ \\ / //  |  _|\"<|  \\| |>   \\| |\\| ||  _|\"      \n\t  | |_) |    | |     | |___U| |\\  |u /\\ V /_,-.| |___U| |\\  |u    | |_| || |___      \n\t  |____/   U/| |\\u   |_____||_| \\_| U  \\_/-(_/ |_____||_| \\_|    <<\\___/ |_____|     \n\t _|| \\_.-,_|___|_,-.<<   >>||   \\,-.//       <<   >>||   \\,-.(__) )(  <<   >>  \n\t(__) (__)\\_)-' '-(_/(__) (__|_\"  (/_/(__)     (__) (__|_\")  (_/     (__)(__) (__)    \n", 10, "cyan")
	fmt.Println("                                                           -                                                        ")
	fmt.Println("                                                           -                                                        ")
	fmt.Println("                                                           -                                                        ")
	SpeedMsg("----------Il y a longtemps, dans un royaume lointain, le maléfique dragon Smaug a volé l'Orbe de la Destinée----------\n", 20, "green")
	SpeedMsg("----------une puissante relique capable de contrôler la magie elle-même. Le royaume est plongé dans les ténèbres\n", 20, "green")
	SpeedMsg("----------et le désespoir règne parmi le peuple.\n", 20, "red")
	SpeedMsg("\n\t       _______________ ______________________ ________________________ \n\t /  _  \\   \\ /   /\\_   _____/ \\      \\__    ___/    |   \\______   \\_   _____/ \n\t/  /_\\  \\   Y   /  |    __)_  /   |   \\|    |  |    |   /|       _/|    __)_\n\t/    |    \\     /   |        \\/    |    \\    |  |    |  / |    |   \\|        \\ \n\t\\____|__  /\\___/   /_______  /\\____|__  /____|  |______/  |____|_  /_______  / \n\t\\/                 \\/         \\/                         \\/        \\/           ", 10, "cyan")
	fmt.Println()
	SpeedMsg("----------Vous êtes un jeune aventurier courageux, prêt à tout pour sauver votre royaume. Vous avez entendu parler----------\n", 20, "green")
	SpeedMsg("----------d'une prophétie ancienne qui dit que seul un héros digne pourra reprendre l'Orbe de la Destinée et----------\n", 20, "green")
	SpeedMsg("----------défaire Smaug. C'est votre destinée qui vous appelle, et vous partez pour une aventure épique.----------\n", 20, "green")
	fmt.Println()
	SpeedMsg("----------Votre voyage commence ici. Choisissez avec sagesse votre personnage et préparez-vous à affronter des----------\n", 20, "green")
	SpeedMsg("----------dangers inimaginables. L'avenir du royaume repose entre vos mains.----------\n", 20, "green")
	fmt.Println()
	SpeedMsg("-------------------------------( Appuyez sur Entrée pour commencer votre quête...) -------------------------------\n", 20, "blue")
	fmt.Scanln() // Attendez l'entrée de l'utilisateur pour passer au jeu.
}

// //////////////////////////////////////////////////////////////////////////////////Lancer combat d'entrainement////////////////////////////
func TrainingFight() {
	joueurPv := 50 // Points de vie du joueur
	tourDeCombat := 1

	gobelinNom := "Gobelin"
	gobelinPvMax := 30
	gobelinPvActuels := gobelinPvMax
	gobelinAttaque := 5

	fmt.Printf("Un %s apparaît !\n", gobelinNom)
	time.Sleep(1000000000)

	for {
		fmt.Printf("\nTour %d de combat\n", tourDeCombat)

		// Tour du joueur
		fmt.Printf("Points de vie du joueur : %d/%d\n", joueurPv, 50)
		fmt.Printf("Points de vie du %s : %d/%d\n", gobelinNom, gobelinPvActuels, gobelinPvMax)

		// Le joueur choisit son attaque
		var choixAttaque int
		fmt.Println("Choisissez votre attaque :")
		fmt.Println("1. Boule de feu")
		fmt.Println("2. Poison")
		fmt.Println("3. Boit ta potion de vie")
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choixAttaque)

		switch choixAttaque {
		case 1:
			fmt.Println("Le joueur lance une boule de feu !")
			gobelinPvActuels -= 15 // Supposons que la boule de feu inflige 15 points de dégâts
		case 2:
			fmt.Println("Le joueur utilise du poison !")
			gobelinPvActuels -= 5 // Supposons que le poison inflige 5 points de dégâts par tour
		case 3:
			fmt.Println("Le joueur boit sa potion de vie !")
			p1.TakePot2()
		default:
			fmt.Println("Attaque non reconnue. Le joueur perd son tour.")
		}

		if gobelinPvActuels <= 0 {
			fmt.Printf("Le joueur a vaincu le %s !\n", gobelinNom)
			break
		}

		// Tour du monstre
		fmt.Println("Le monstre attaque le joueur !")
		joueurPv -= gobelinAttaque // Le monstre inflige des dégâts égaux à son attaque

		if joueurPv <= 0 {
			fmt.Printf("%s a vaincu le joueur. Le combat est terminé.\n", gobelinNom)
			break
		}
		// Code pour le combat avec le gobelin...

		VictoryAnimation() // Affichez l'animation de victoire

		// Continuez avec le reste du combat ou retournez au menu principal, selon votre logique de jeu.

		tourDeCombat++
	}
}

// ///////////////////////////////////////////////////////////////////////////////GROS GOBELIN///////////////////////////////////////////////
func InitGoblin() Monstre {
	gobelin := Monstre{
		Name:      "Gobelin d'entraînement",
		MaxHealth: 40,
		Health:    40, // Les points de vie actuels sont égaux aux points de vie maximum au début
		Attack:    5,
	}
	return gobelin

}

// ////////////////////////////////////////////////////////////tous les perso qui exisent
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
	RaconterHistoire()
	StartScreen()
	// Crée un gobelin d'entraînement en utilisant la fonction InitGoblin
	gobelin := InitGoblin()

	// Affiche les caractéristiques du gobelin
	fmt.Printf("Un %s apparaît !\n", gobelin.Name)
	fmt.Printf("Points de vie : %d/%d\n", gobelin.Health, gobelin.MaxHealth)
	fmt.Printf("Points d'attaque : %d\n", gobelin.Attack)

	// Supposons que vous ayez déjà créé un joueur et un Gobelin avec InitGoblin
	player := &Character{
		Name:      "Personnage",
		Health:    100,
		MaxHealth: 100,
		Attack:    15,
	}

	// Appel de la fonction goblinPattern
	GoblinPattern(player, &gobelin)

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

func VictoryAnimation() {
	// Vous pouvez utiliser des chaînes ASCII art pour créer votre animation.
	animation := `
    __   __            _    _ _       
    \ \ / /__ __ _ __| |__| | | ___  
     \ V / -_| '_/ _| '_ \ | |(_-<  
      \_/\___|_| \__|_.__/_|_|/__/
    `
	fmt.Println(animation)
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

func (m *Monstre) InitGobbelin(nom string, pvmax int, pv int, pvAttack int) {

	m.Name = nom
	m.MaxHealth = pvmax
	m.Health = pv
	m.Attack = pvAttack

}

func (p *Monstre) TrainingFight() {
	var Gobelin Monstre
	Gobelin.InitGobbelin("Serge le Gobelin", 40, 40, 10)

	///////////////////////////////////////////////////////////FAIRE UN MARCHAND EN ASCII

	///////////////////////////////////////////////////////////FAIRE UN FORGERON EN ASCII

	///////////////////////////////////////////////////////////RAJOUTER DES ANIMATIONS LORS DU COMBATS
}
