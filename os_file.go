package piscine

import (
	"fmt"
	"os"
)

func open_file(file string) {
	fichier, err := os.Open(file)
	if err != nil {
		fmt.Println("erreur d'ouverture", err)
		return
	}
	defer fichier.Close()
}

func return_file(text string) string {
	os.WriteFile("correction.txt", []byte(text), 0644)
	return "version du fichier corrigée envoyée"
}
