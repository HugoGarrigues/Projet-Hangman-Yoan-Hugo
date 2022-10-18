package Dossier_Package

import "fmt"

func (j *Joueur) lancement() {
	mot := motAleatoire()
	mot_cache := masquerMot(mot)
	for !motEstTrouve(mot_cache) && j.essais != 0 {
		fmt.Println("Les lettres que vous avez essayé sont :", j.chaine_mot)
		afficheMotCache(mot_cache)
		lettre := saisieLettre()
		if lettreEstPresente(lettre, mot) {
			mot_cache = afficheMotAvecLettreTrouvee(lettre, mot, mot_cache)
			fmt.Print("Bravo ! La lettre ", lettre, " est présente dans le mot\n")
		} else {
			j.essais--
			fmt.Print("Dommage ! La lettre ", lettre, " n'est pas présente dans le mot\n Il vous reste ", j.essais, " essais\n")
		}
	}
	afficheResultat(mot_cache, mot, j.essais)
}
