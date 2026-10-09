package piscine

import (
	"fmt"
	"os"
)

func open_file(file string) error {
	fichier, err := os.Open(file)
	if err != nil {
		fmt.Println("erreur d'ouverture", err)
		return err
	}
	defer fichier.Close()
	return nil
}

func return_file(text string) string {
	os.WriteFile("correction.txt", []byte(text), 0644)
	return "version du fichier corrigée envoyée"
}
