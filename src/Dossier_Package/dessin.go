package Dossier_Package

import (
	"os"
	"bufio"
	"log"
	"fmt"
)

func (j *Joueur)dessin_erreur(){
	switch j.essais{
	case 9:
		betweenWord("z", "a")
	case 8:
		betweenWord("a", "b")
	case 7:
		betweenWord("b", "c")
	case 6:
		betweenWord("c", "d")
	case 5:
		betweenWord("d", "e")
	case 4:
		betweenWord("e", "f")
	case 3:
		betweenWord("f", "g")
	case 2:
		betweenWord("g", "h")
	case 1:
		betweenWord("h", "i")
	case 0:
		betweenWord("i", "j")
	}			
}	




func betweenWord(mot1 string, mot2 string) {
	file, err := os.Open("Hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	var t bool
	for scan.Scan() {
		nline := scan.Text()
		if !t {
			t = nline == mot1
		}
		if t {
			if nline == mot2 {
				break
			}
			if nline != mot1 {
				fmt.Println(nline)
			}
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
}
