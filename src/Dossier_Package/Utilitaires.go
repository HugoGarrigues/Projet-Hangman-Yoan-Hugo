package Dossier_Package

import (
	"strconv"
)

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

func contains(slice []int, element int) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
