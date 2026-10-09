package main

import (
	"fmt"
	"strings"
)

func up(strup string) string {
	return fmt.Sprintf("%s", strings.ToUpper(strup))
}

func low(strlow string) string {
	return fmt.Sprintf("%s", strings.ToLower(strlow))
}

func cap(strcap string) string {
	return fmt.Sprintf("%s", strings.Title(strcap))
}
