package main

import (
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	dir := t.TempDir()
	chemin := dir + "/test.txt"
	os.WriteFile(chemin, []byte("test de la fonction"), 0644)
	err := open_file(chemin)
	if err != nil {
		t.Error(err)
	}
}
