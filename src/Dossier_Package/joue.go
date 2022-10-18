package Dossier_Package

import "fmt"

func (j *Joueur) lancement() {
	mot := motAleatoire()
	mot_cache := masquerMot(mot)
	nbEssais := 10
	for nbEssais > 0 {
		fmt.Println("Il vous reste ", nbEssais, " essais")
		afficheMotCache(mot_cache)
		lettre := saisieLettre()
		if verifierLettre(lettre, mot) {
			mot_cache = afficheMotAvecLettreTrouvee(lettre, mot, mot_cache)
		} else {
			nbEssais--
		}
	}
	afficheResultat(mot_cache, mot, nbEssais)
}
