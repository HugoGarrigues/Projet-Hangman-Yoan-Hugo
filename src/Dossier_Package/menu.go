package Dossier_Package

import (
	"strings"
)

func (j *Joueur) menu() {
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
