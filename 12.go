package main

import "fmt"

// Escreva um programa que receba uma string e
//verifique se ela é um palíndromo.
//Informe ao usuário se é ou não.

func main() {
	var x string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&x)

	var reverse string

	for i := len(x) - 1; i >= 0; i-- {
		reverse = reverse + string(x[i])
	}

	if x == reverse {
		fmt.Println("É palindromo")
	} else {
		fmt.Println("Não é palindromo")
	}
}
