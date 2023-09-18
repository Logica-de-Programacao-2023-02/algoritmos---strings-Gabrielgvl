package main

import "fmt"

// Escreva um programa que receba duas strings e
// verifique se elas são anagramas.
// Informe ao usuário se são ou não.

func main() {
	var x, y string
	fmt.Print("Escreva duas palavras: ")
	fmt.Scan(&x, &y)

	xMapa := make(map[string]int)
	yMapa := make(map[string]int)

	for _, c := range x {
		xMapa[string(c)]++
	}

	for _, c := range y {
		yMapa[string(c)]++
	}
	
	for char, qtdX := range xMapa {
		qtdY := yMapa[char]
		if qtdX != qtdY {
			fmt.Println("As palavras não são anagramas")
			return
		}
	}

	fmt.Println(xMapa)
	fmt.Println(yMapa)

	fmt.Println("As palavras são anagramas")
}
