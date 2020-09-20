package main

import (
	"fmt"
	"testing"
)

func Soma(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Soma(5, 5))
}

func TestSoma(t* testing.T) {
	result = Soma(2, 2)

	if result != 4{
		t.Errorf("Soma espera: %d, obitida: %d", 4, result)
	} else {
		fmt.Println("Passou")
	}
}