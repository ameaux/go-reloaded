package piscine

import (
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	dir := t.TempDir()
	chemin := dir + "/test.txt"
	os.WriteFile(chemin, []byte("test de la fonction"), 0644)
	open_file(chemin)
}
