package main

import "fmt"

// Escreva um programa que receba uma string e
//inverta a ordem dos caracteres.
//Informe ao usuário o resultado.

func main() {
	var x string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&x)

	var reverse string

	for i := len(x) - 1; i >= 0; i-- {
		// "" + "i"
		// "i" + "x"
		// "ix" + "a"
		// ...
		// "ixacaba"
		reverse = reverse + string(x[i])
	}

	fmt.Println(reverse)
}
