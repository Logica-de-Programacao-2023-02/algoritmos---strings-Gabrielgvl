package main

import (
	"fmt"
)

// Escreva um programa que receba uma string e
//remova todas as vogais.
//Informe ao usuário o resultado.

func main() {
	var x string

	fmt.Print("Escreva uma string: ")
	fmt.Scan(&x)

	vogais := []string{
		"A",
		"a",
		"E",
		"e",
		"I",
		"i",
		"O",
		"o",
		"U",
		"u",
	}
	var result string
	for _, c := range x {
		var isVogal bool
		for _, vogal := range vogais {
			if string(c) == vogal {
				isVogal = true
				break
			}
		}
		if !isVogal {
			result += string(c)
		}
	}

	fmt.Println(result)
}
