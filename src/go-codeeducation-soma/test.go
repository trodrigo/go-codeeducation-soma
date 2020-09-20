package main

import (
	"testing"
)

func Soma(a, b int) int {
	return a + b
}

func TestSoma(t* testing.T) {
	result = Soma(2, 2)

	if result != 4{
		t.Errorf("Soma espera: %d, obitida: %d", 4, result)
	}
}