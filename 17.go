package main

import "fmt"

// Solicite ao usuário uma string e imprima somente
//as suas letras únicas (que aparecem apenas uma vez).

func main() {
	var x string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&x)

	var unicas []string

	for i := 0; i < len(x); i++ {
		isUnique := true
		for j := 0; j < len(x); j++ {
			if i != j && x[i] == x[j] {
				isUnique = false
				break
			}
		}
		if isUnique {
			unicas = append(unicas, string(x[i]))
		}
	}

	fmt.Println(unicas)
}
