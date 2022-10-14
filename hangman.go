package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	compteur := 0
	words := chooseAleatoryWords()
	var letter string
	var choix string
	for {
		if compteur == 0 {
			fmt.Println("Bienvenue dans le jeu du Pendu")
		}
		compteur++
		if compteur == 10 {
			fmt.Println("You lost")
			break
		}
		fmt.Println("Mot à deviner : ")
		revealWord(words)
		fmt.Println("\n")
		fmt.Println("")
		fmt.Println("Il vous reste ", 10-compteur, " essais")
		fmt.Println("1. Deviner une lettre")
		fmt.Println("2. Deviner le mot")
		fmt.Scanln(&choix)
		switch choix {
		case "1":
			fmt.Println("Devinez une lettre")
			fmt.Scanln(&letter)
			if len(letter) > 1 {
				fmt.Println("Vous devez entrer une seule lettre")
			} else {
				checkLetter(words, letter)
			}
		case "2":
			fmt.Println("Devinez le mot")
			fmt.Scanln(&letter)
			if letter == words {
				fmt.Println("You won")
				break
			} else {
				fmt.Println("You lost")
				break
			}
		default:
			fmt.Println("Veuiilez entrer 1 ou 2")
		}
	}
}

func chooseAleatoryWords() string {
	words := []string{}
	file, err := os.Open("hangman_words.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

/* The programm will reveal n random letters in the word, where n is the len(word) / 2 - 1 */

func revealWord(mot string) {
	nbLettre := len(mot)/2 - 1
	if nbLettre < 1 {
		for i := 0; i < len(mot); i++ {
			fmt.Print("_ ")
		}
	} else {
		for i := 0; i < nbLettre; i++ {
			fmt.Print("_ ")
		}
		for i := nbLettre; i < len(mot); i++ {
			fmt.Print(string(mot[i]), " ")
		}
	}
}

func checkWord(mot string, letter string) bool {
	for i := 0; i < len(mot); i++ {
		if string(mot[i]) == letter {
			return true
		}
	}
	return false
}

func checkLetter(mot string, letter string) {
	if checkWord(mot, letter) {
		fmt.Println("Correct")
		actualiseWord(mot, letter)
	} else {
		fmt.Println("Wrong")
	}
}

func actualiseWord(mot string, letter string) {
	for i := 0; i < len(mot); i++ {
		if string(mot[i]) == letter {
			fmt.Println(mot[i])
		} else {
			fmt.Printf("_ ")
		}
	}
}
