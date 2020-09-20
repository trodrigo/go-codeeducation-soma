package main

import (
	"testing"
	"fmt"
)

func TestSoma(t* testing.T) {
	result = Soma(2, 2)

	if result != 4{
		t.Errorf("Soma espera: %d, obitida: %d", 4, result)
	} else {
		fmt.Println("Passou")
	}
}