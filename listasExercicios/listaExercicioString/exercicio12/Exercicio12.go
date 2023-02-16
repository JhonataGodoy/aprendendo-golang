package main

//Valida e corrige número de telefone. Faça um programa que leia um número de telefone, e corrija
//o número no caso deste conter somente 7 dígitos, acrescentando o '3' na frente. O usuário pode
//informar o número com ou sem o traço separador.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your phone number:")
	pickUpNumber, _ := reader.ReadString('\n')
	wipeNumber := strings.TrimSpace(pickUpNumber)
	CaracterNumber := len(wipeNumber) - 1

	if CaracterNumber == 7 {
		wipeNumber = string(wipeNumber) + "3"
		fmt.Print("telefone:", pickUpNumber)
		fmt.Println("Telefone possui 7 dígitos. Vou acrescentar o digito três na frente.")
		fmt.Print("Telefone corrigido sem formatação:", pickUpNumber)
		fmt.Print("Telefone corrigido com formatação:", wipeNumber)
	}

	if CaracterNumber == 8 || CaracterNumber == 9 {
		fmt.Print("telefone:", pickUpNumber)
	}

}

//Valida e corrige número de telefone
//Telefone: 461-0133
//Telefone possui 7 dígitos. Vou acrescentar o digito três na frente.
//Telefone corrigido sem formatação: 34610133
//Telefone corrigido com formatação: 3461-0133
