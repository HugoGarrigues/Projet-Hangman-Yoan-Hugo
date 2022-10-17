package Dossier_Package

import (
	"bufio"
	"fmt"
	"os"
)

type Joueur struct{
	nom string
	essais int
	chaine_mot []string
}

func (j *Joueur) Initialisation() {
	for {
		fmt.Println("Choisissez votre nom comprenant que des lettres (pas d'espace ni d'accent):")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			j.nom = line
		}
		if j.verif_nom(j.nom) {
			j.nom = j.majuscule(j.nom)
			fmt.Println("Votre nom est ", j.nom)
			j.essais = 10
			break
		}
		fmt.Println("Votre nom est invalide")
	}
}