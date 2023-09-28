package red

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

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

// /////////////////////////////////////////////////////////////////////////FONCTION CLEAR CONSOLE
func ClearConsole() {
	switch runtime.GOOS {
	case "darwin":
		RunCmd("clear")
	case "linux":
		RunCmd("clear")
	case "windows":
		RunCmd("cmd", "/c", "cls")
	default:
		RunCmd("clear")
	}
}

// ///////////////////////////////////////////////////////////////////////////CLearConsole efface La console.
func RunCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
