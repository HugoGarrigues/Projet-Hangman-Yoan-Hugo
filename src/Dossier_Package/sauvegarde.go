package Dossier_Package

import (
	"bufio"
	"fmt"
	"os"
)

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

func (j *Joueur) RecupérationDonnéesSauvegarde() {
	var nom string
	fmt.Println("Saisissez le nom utilisé lors de votre ancienne partie :")
	fmt.Scan(&nom)
	var fichier string
	fichier = nom + ".txt"
	f, err := os.Open(fichier)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var mot_cache string
	var mot string
	var chaine_mot string
	var nom_joueur string
	var essais string
	for scanner.Scan() {
		mot_cache = scanner.Text()
		scanner.Scan()
		mot = scanner.Text()
		scanner.Scan()
		chaine_mot = scanner.Text()
		scanner.Scan()
		nom_joueur = scanner.Text()
		scanner.Scan()
		essais = scanner.Text()
	}
	essaisInt := IntInString(essais)
	var choix string
	fmt.Println("1. Continuer la partie")
	fmt.Println("2. Supprimer la partie")
	fmt.Scan(&choix)
	if choix == "1" {
		j.nom = nom_joueur
		j.chaine_mot = chaine_mot
		j.essais = essaisInt
		j.ContinuerPartie(mot_cache, mot)
	} else if choix == "2" {
		os.Remove(fichier)
	} else {
		fmt.Println("Veuillez saisir un choix valide.")
	}
}
