package Dossier_Package

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Joueur struct {
	nom        string
	essais     int
	chaine_mot string
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

func (j *Joueur) Sauvegarder1(mot_cache string, mot string, essais int, chaine_mot string, nom string) {
	var fichier string
	essaisString := stringInInt(essais)
	fichier = nom + ".txt"
	f, err := os.Create(fichier)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.WriteString(mot_cache)
	f.WriteString("\n")
	f.WriteString(mot)
	f.WriteString("\n")
	f.WriteString(chaine_mot)
	f.WriteString("\n")
	f.WriteString(nom)
	f.WriteString("\n")
	f.WriteString(essaisString)
}

func stringInInt(essais int) string {
	var chaine_mot string
	chaine_mot = strconv.Itoa(essais)
	return chaine_mot
}

func IntInString(chaine_mot string) int {
	var essais int
	essais, _ = strconv.Atoi(chaine_mot)
	return essais
}
