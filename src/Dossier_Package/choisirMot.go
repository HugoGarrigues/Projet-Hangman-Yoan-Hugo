package Dossier_Package

import (
	"bufio"
	"fmt"
	"os"
)

func ChoisirMot() string {
	file, err := os.Open("Hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		break
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	file.Close()
	mot := fileScanner.Text()
	return mot
}
