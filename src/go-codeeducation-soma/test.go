package main

import (
	"testing"
)

func Soma2(a, b int) int {
	return a + b
}

func TestSoma(t* testing.T) {
	result = Soma2(2, 2)

	if result != 4{
		t.Errorf("Soma espera: %d, obitida: %d", 4, result)
	}
}