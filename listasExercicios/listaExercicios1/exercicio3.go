package main

//Faça um Programa que peça dois números e imprima a soma.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pegaN1 := 0
	pegaN2 := 0
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(" Digite um numero: ")
		pegaNumero1, _ := reader.ReadString('\n')
		limpaNumero1 := strings.TrimSpace(pegaNumero1)
		numero1, err := strconv.Atoi(limpaNumero1)
		if err != nil {
			fmt.Println("O que você digitou não é um numero ")
			continue
		}
		pegaN1 = numero1
		break
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(" Digite um numero: ")
		pegaNumero2, _ := reader.ReadString('\n')
		limpaNumero2 := strings.TrimSpace(pegaNumero2)
		numero2, err := strconv.Atoi(limpaNumero2)
		if err != nil {
			fmt.Println("O que você digitou não é um numero ")
			continue
		}
		pegaN2 = numero2
		break
	}

	fmt.Println("Soma dos dois numeros digitados: ", pegaN1+pegaN2)

}
