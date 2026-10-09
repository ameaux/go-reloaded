package piscine

import (
	"fmt"
	"strconv"
)

func convert_base_hex(numex string) string {
	number, err := strconv.Atoi(numex)
	if err != nil {
		fmt.Println("erreur de conversion", err)
		return ""
	}
	return fmt.Sprintf("%d", number)
}

func convert_base_bin(numin string) string {
	nombre, err := strconv.ParseInt(numin, 2, 64)
	if err != nil {
		fmt.Println("erreur de conversion :", err)
		return ""
	}
	return fmt.Sprintf("%d", nombre)
}
