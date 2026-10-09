package main

import (
	"fmt"
	"os"
	"strings"
)

func parcours(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	sep := strings.Split(string(content), " ")
	for i := 0; i < len(sep); i++ {
		if sep[i] == "(hex)" {
			sep[i-1] = convert_base_hex(sep[i-1])
		}
		if sep[i] == "(bin)" {
			sep[i-1] = convert_base_bin(sep[i-1])
		}
		if sep[i] == "(up)" {
			sep[i-1] = up(sep[i-1])
		}
		if sep[i] == "(low)" {
			sep[i-1] = low(sep[i-1])
		}
		if sep[i] == "(cap)" {
			sep[i-1] = cap(sep[i-1])
		}
	}
	return strings.Join(sep, " ")
}
