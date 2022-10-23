package Dossier_Package

import (
	"os"
	"bufio"
	"log"
	"fmt"
)

func (j *Joueur)dessin_erreur(){
	if j.essais == 9 {
		betweenWord("z", "a")
	}else if j.essais == 8 {
		betweenWord("a", "b")
	}else if j.essais == 7 {
		betweenWord("b", "c")
	}else if j.essais == 6 {
		betweenWord("c", "d")
	}else if j.essais == 5 {
		betweenWord("d", "e")
	}else if j.essais == 4 {
		betweenWord("e", "f")
	}else if j.essais == 3 {
		betweenWord("f", "g")
	}else if j.essais == 2 {
		betweenWord("g", "h")
	}else if j.essais == 1 {
		betweenWord("h", "i")
	}else if j.essais == 0 {
		betweenWord("i", "j")
	}			
}	




func betweenWord(mot1 string, mot2 string) {
	var betweenWord string
	nmot := ""
	file, err := os.Open("Hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		nmot = scan.Text()
		if nmot == mot1 {
			for scan.Scan() {
				betweenWord = scan.Text()
				if betweenWord == mot2 {
					break
				}
				fmt.Println(betweenWord)
			}
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
}