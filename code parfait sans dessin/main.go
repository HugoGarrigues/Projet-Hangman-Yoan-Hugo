package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func choisirMot() string {
	f, err := os.Open("mot.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var mots []string
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}
	rand.Seed(time.Now().UnixNano())
	return mots[rand.Intn(len(mots))]
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

// The program will read the standard input to suggest a letter. //

func lireLettre() string {
	var lettre string
	fmt.Println("Choose a letter: ")
	fmt.Scan(&lettre)
	return lettre
}

func verifierLettre(lettre string, mot string) bool {
	for _, l := range mot {
		if string(l) == lettre {
			return true
		}
	}
	return false
}

func remplacerLettre(lettre string, mot string, motMasque string) string {
	runes := []rune(mot)
	runesMasque := []rune(motMasque)
	for i, l := range runes {
		if string(l) == lettre {
			runesMasque[i] = l
		}
	}
	return string(runesMasque)
}

func verifierMot(motMasque string) bool {
	for _, l := range motMasque {
		if l == '_' {
			return false
		}
	}
	return true
}

func main() {
	mot := choisirMot()
	motMasque := masquerMot(mot)
	fmt.Println(motMasque)
	for {
		lettre := lireLettre()
		if verifierLettre(lettre, mot) {
			motMasque = remplacerLettre(lettre, mot, motMasque)
			fmt.Println(motMasque)
		}
		if verifierMot(motMasque) {
			fmt.Println("You win!")
			break
		}
	}
}
