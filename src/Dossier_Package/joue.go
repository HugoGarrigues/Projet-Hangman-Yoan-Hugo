package Dossier_Package

import (
	"fmt"
)

func (j *Joueur) PremierLancement() {
	mot := motAleatoire()
	mot_cache := masquerMot(mot)
	for !motEstTrouve(mot_cache) && j.essais != 0 {
		fmt.Println("Mot à trouver : ", mot_cache)
		var choix string
		fmt.Println("1. Saisir une lettre")
		fmt.Println("2. Saisir un mot")
		fmt.Println("3. Sauvegarder et quitter")
		fmt.Scan(&choix)
		if choix == "1" {
			fmt.Println("Les lettres que vous avez essayé sont :", j.chaine_mot)
			lettre := saisieLettre()
			if j.ajoutLettre(lettre) == false {
				fmt.Println("Vous avez déjà essayé cette lettre veuillé en saisir une autre, il vous reste ", j.essais, " essais.")
			} else {
				j.ajoutLettre(lettre)
				if lettreEstPresente(lettre, mot) {
					mot_cache = afficheMotAvecLettreTrouvee(lettre, mot, mot_cache)
					fmt.Print("Bravo ! La lettre ", lettre, " est présente dans le mot\n")
				} else {
					j.essais--
					fmt.Print("Dommage ! La lettre ", lettre, " n'est pas présente dans le mot\n Il vous reste ", j.essais, " essais\n")
					j.dessin_erreur()
				}
			}
		} else if choix == "2" {
			var mot_saisi string
			fmt.Println("Saisir un mot : ")
			fmt.Scan(&mot_saisi)
			if mot_saisi == mot {
				fmt.Println("Bravo ! Vous avez trouvé le mot !")
				break
			} else {
				j.essais -= 2
				fmt.Println("Dommage ! Le mot n'est pas correct, il vous reste ", j.essais, " essais.")
			}
		} else if choix == "3" {
			j.Sauvegarder1(mot_cache, mot, j.essais, j.chaine_mot, j.nom)
			break
		} else {
			fmt.Println("Veuillez saisir un choix valide.")
		}
	}
	if j.essais == 0 {
		afficheResultat(mot_cache, mot, j.essais)
	}
}

func (j *Joueur) ContinuerPartie(mot_cache string, mot string) {
	for !motEstTrouve(mot_cache) && j.essais != 0 {
		fmt.Println("Mot à trouver : ", mot_cache)
		var choix string
		fmt.Println("1. Saisir une lettre")
		fmt.Println("2. Saisir un mot")
		fmt.Println("3. Sauvegarder et quitter")
		fmt.Scan(&choix)
		if choix == "1" {
			fmt.Println("Les lettres que vous avez essayé sont :", j.chaine_mot)
			lettre := saisieLettre()
			if j.ajoutLettre(lettre) == false {
				fmt.Println("Vous avez déjà essayé cette lettre veuillé en saisir une autre, il vous reste ", j.essais, " essais.")
			} else {
				j.ajoutLettre(lettre)
				if lettreEstPresente(lettre, mot) {
					mot_cache = afficheMotAvecLettreTrouvee(lettre, mot, mot_cache)
					fmt.Print("Bravo ! La lettre ", lettre, " est présente dans le mot\n")
				} else {
					j.essais -= 2
					fmt.Print("Dommage ! La lettre ", lettre, " n'est pas présente dans le mot\n Il vous reste ", j.essais, " essais\n")
				}
			}
		} else if choix == "2" {
			var mot_saisi string
			fmt.Println("Saisir un mot : ")
			fmt.Scan(&mot_saisi)
			if mot_saisi == mot {
				fmt.Println("Bravo ! Vous avez trouvé le mot !")
				break
			} else {
				j.essais -= 2
				fmt.Println("Dommage ! Le mot n'est pas correct, il vous reste ", j.essais, " essais.")
			}
		} else if choix == "3" {
			j.Sauvegarder1(mot_cache, mot, j.essais, j.chaine_mot, j.nom)
			break
		} else {
			fmt.Println("Veuillez saisir un choix valide.")
		}
	}
	if j.essais == 0 || motEstTrouve(mot_cache) {
		afficheResultat(mot_cache, mot, j.essais)
	}
}
