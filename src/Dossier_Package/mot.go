package Dossier_Package

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func motAleatoire() string {
	fichier, err := os.Open("mot.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier !")
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)
	var mots []string
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
	return mots[rand.Intn(len(mots))]
}

func afficheMotCache(motCache string) {
	fmt.Println(motCache)
}

func saisieLettre() string {
	var lettre string
	fmt.Println("Saisir une lettre : ")
	fmt.Scan(&lettre)
	if len(lettre) != 1 {
		fmt.Println("Vous devez saisir une seule lettre ! ")
		return saisieLettre()
	} else if lettre == " " {
		fmt.Println("Vous devez saisir une lettre ! ")
		return saisieLettre()
	} else if lettre == "" {
		fmt.Println("Vous devez saisir une lettre ! ")
		return saisieLettre()
	} else if verif_lettre(lettre) == false {
		fmt.Println("Vous devez saisir une lettre ! ")
		return saisieLettre()
	} else if verif_minuscule(lettre) == false {
		lettre = minuscule(lettre)
	}
	return lettre
}

func afficheMotAvecLettreTrouvee(lettre string, mot string, motCache string) string {
	var motCacheTemporaire string
	for i := 0; i < len(mot); i++ {
		if string(mot[i]) == lettre {
			motCacheTemporaire += lettre
		} else {
			motCacheTemporaire += string(motCache[i])
		}
	}
	return motCacheTemporaire
}

func afficheResultat(motCache string, mot string, nombreEssais int) {
	if motEstTrouve(motCache) {
		fmt.Println("Bravo, vous avez trouvé le mot ! \n Le mot etait ", mot)
	} else {
		fmt.Println("Vous avez perdu ! Le mot était : ", mot)
	}
}

func masquerMot(mot string) string {
	n := len(mot)/2 - 1
	runes := []rune(mot)
	var indexes []int
	var revealed []rune
	for i := 0; i < n; i++ {
		index := rand.Intn(len(runes))
		for contains(indexes, index) {
			index = rand.Intn(len(runes))
		}
		indexes = append(indexes, index)
		revealed = append(revealed, runes[index])
	}
	for i := 0; i < len(runes); i++ {
		if !contains(indexes, i) {
			runes[i] = '_'
		}
	}

	return string(runes)
}

func contains(slice []int, element int) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func (j *Joueur) ajoutLettre(lettre string) {
	if lettreEstPresente(lettre, j.chaine_mot) {
		fmt.Println("Vous avez déjà essayé cette lettre !")
		j.essais++
	} else {
		j.chaine_mot += lettre
	}
}
