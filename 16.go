package main

import (
	"fmt"
	"strings"
)

// Solicite ao usuário duas strings e
//informe se a segunda é uma substring da primeira.

func main() {
	var x, y string
	fmt.Print("Escreva a string principal: ")
	fmt.Scan(&x)

	fmt.Print("Escreva a string que está contida: ")
	fmt.Scan(&y)

	if strings.Contains(x, y) {
		fmt.Println("A segunda é uma substring da primeira")
	} else {
		fmt.Println("Tem nada a ver")
	}
}
