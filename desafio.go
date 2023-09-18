package main

import (
	"fmt"
	"strings"
)

// Escreva um programa que receba uma string
// e um padrão (outro string) e
// retorne todos os índices em que o
// padrão ocorre na string.
// Informe ao usuário o resultado.

func main() {
	var x, padrao string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&x)
	fmt.Print("Digite um padrao: ")
	fmt.Scan(&padrao)

	var indexes []int

	for i := len(x) - 1; i >= 0; i -= len(padrao) {
		idx := strings.Index(x[i:], padrao)
		if idx != -1 {
			indexes = append(indexes, i+idx)
		}
	}

	fmt.Println(indexes)
}
