package main

import "fmt"

func main() {
	fmt.Println(Ola(""))
}

const prefixoOlaPortugues = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixoOlaPortugues + nome
}
