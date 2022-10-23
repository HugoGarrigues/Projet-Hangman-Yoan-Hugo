package Dossier_Package

import (
	"fmt"
	"os"
)

func (j *Joueur) Menu() {
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|                                     |")
	fmt.Println("|    BIENVENUE A VOUS DANS CE JEU     |")
	fmt.Println("|                                     |")
	fmt.Println("|             Le Pendu                |")
	fmt.Println("|                                     |")
	fmt.Println("|         1. Commencer le jeu         |")
	fmt.Println("|         2. Commencer le jeu a       |")
	fmt.Println("|            partir d'une sauvegarde  |")
	fmt.Println("|         3. Quitter le jeu           |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		j.Initialisation()
		j.PremierLancement()
	case 2:
		j.RecupérationDonnéesSauvegarde()
	case 3:
		fmt.Println("Merci d'avoir joué")
		os.Exit(0)
	default:
		fmt.Println("Choix invalide")
		j.Menu()
	}
}
