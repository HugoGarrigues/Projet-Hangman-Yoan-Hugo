package Dossier_Package

import (
	"strings"
	"unicode"
)

func (j *Joueur) majuscule(nom string) string {
	var nom_maj string
	for i, letter := range nom {
		if i == 0 {
			nom_maj += string(unicode.ToUpper(letter))
		} else {
			nom_maj += string(unicode.ToLower(letter))
		}
	}
	return nom_maj
}

func (j *Joueur) verif_nom(nom string) bool {
	var nb_lettre int
	for _, letter := range nom {
		if unicode.IsLetter(letter) {
			nb_lettre += 1
		}
	}
	return nb_lettre == len(nom)
}

func lettreEstPresente(lettre string, mot string) bool {
	return strings.Contains(mot, lettre)
}

func motEstTrouve(motCache string) bool {
	return !strings.Contains(motCache, "_")
}

func nombreEssaisEpuise(nombreEssais int) bool {
	return nombreEssais == 0
}
