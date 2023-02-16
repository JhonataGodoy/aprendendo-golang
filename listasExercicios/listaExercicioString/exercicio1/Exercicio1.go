package main

//Tamanho de strings. Faça um programa que leia 2 strings e informe o conteúdo delas seguido do seu comprimento.
//Informe também se as duas strings possuem o mesmo comprimento e são iguais ou diferentes no conteúdo.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var primeiraString string
	var segundaString string
	for i := 0; i < 2; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(" Type a string: ")
		pegaString, _ := reader.ReadString('\n')
		limpaString := strings.TrimSpace(pegaString)
		if i == 0 {
			primeiraString = limpaString
		} else {
			segundaString = limpaString
		}
	}

	fmt.Println("Size of", primeiraString, ":", len(primeiraString), "characters")
	fmt.Println("Size of", segundaString, ":", len(segundaString), "characters")

	if len(primeiraString) != len(segundaString) {
		fmt.Println("Strings of different length")
	} else {
		fmt.Println("Strings of the same length")
	}

	if primeiraString == segundaString {
		fmt.Println("Strings with the same content")
	} else {
		fmt.Println("Strings with different contents")
	}
}
