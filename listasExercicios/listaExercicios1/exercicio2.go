package main

//Faça um Programa que peça um número e então mostre a mensagem O número informado foi [número].

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" Digite um numero: ")
	pegaNumero, _ := reader.ReadString('\n')
	limpaNumero := strings.TrimSpace(pegaNumero)
	numero, err := strconv.ParseFloat(limpaNumero, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um numero ")
	}

	fmt.Println("Numero digitado: ", numero)

}
