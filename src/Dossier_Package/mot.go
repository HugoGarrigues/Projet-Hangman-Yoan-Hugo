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
		fmt.Println("Erreur lors de l'ouverture du fichier")
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
		fmt.Println("Bravo, vous avez trouvé le mot !")
	} else {
		fmt.Println("Vous avez perdu ! Le mot était : ", mot)
	}
}

func motCache(mot string) string {
	var motCache string
	nblettreCache := len(mot)/2 - 1
	for i := 0; i < len(mot); i++ {
		if i < nblettreCache {
			motCache += "_"
		} else {
			motCache += string(mot[i])
		}
	}
	return motCache
}
