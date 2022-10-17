package Dossier_Package

import (
	"fmt"
	"os"
	"strings"
)

func (j *Joueur) Menu() {
	var choix int
	fmt.Println("______________________________________")
	fmt.Println("|                                     |")
	fmt.Println("|    BIENVENUE A VOUS DANS CE JEU     |")
	fmt.Println("|                         	           |")
	fmt.Println("|             Le Pendu                |")
	fmt.Println("|                                     |")
	fmt.Println("|         1. Commencer le jeu         |")
	fmt.Println("|         2. Quitter le jeu           |")
	fmt.Println("|_____________________________________|")
	fmt.Println("")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		j.Initialisation()
		j.sousMenu()
	case 2:
		fmt.Println("Merci d'avoir joué")
		os.Exit(0)
	default:
		fmt.Println("Choix invalide")
		j.Menu()
	}
	for {
		j.sousMenu()
	}
}

func (j *Joueur) sousMenu() {
	mot := motAleatoire()
	motCache := strings.Repeat("_", len(mot))
	nombreEssais := 10
	for !motEstTrouve(motCache) && !nombreEssaisEpuise(nombreEssais) {
		afficheMotCache(motCache)
		lettre := saisieLettre()
		if lettreEstPresente(lettre, mot) {
			motCache = afficheMotAvecLettreTrouvee(lettre, mot, motCache)
		} else {
			nombreEssais--
		}
	}
	afficheResultat(motCache, mot, nombreEssais)
}
