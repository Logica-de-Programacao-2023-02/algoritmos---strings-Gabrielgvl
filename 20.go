package main

import (
	"fmt"
	"strings"
)

// Solicite ao usuário uma string e informe se ela é
//está em camelCase e
//quantas palavras possuí.
//CamelCase é quando a primeira letra de
//cada palavra é maiúscula,
//e as demais são minúsculas.
//Exemplo: "estaStringEstaEmCamelCase".

func main() {
	var x string
	fmt.Print("Escreva uma string: ")
	fmt.Scan(&x)

	count := 1

	for _, c := range x {
		if strings.ToUpper(string(c)) == string(c) {
			count++
		}
	}

	fmt.Println(count)
}
