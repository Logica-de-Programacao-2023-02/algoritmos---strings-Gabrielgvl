package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Escreva um programa que receba uma string e
//conte quantas palavras ela contém.
//Informe ao usuário o resultado.

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	frase := scanner.Text()

	palavras := strings.Fields(frase)

	fmt.Printf("Essa frase tem %d palavras", len(palavras))
}
