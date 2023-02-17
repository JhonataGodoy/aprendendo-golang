package main

//Valida e corrige número de telefone. Faça um programa que leia um número de telefone, e corrija
//o número no caso deste conter somente 7 dígitos, acrescentando o '3' na frente. O usuário pode
//informar o número com ou sem o traço separador.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your phone number:")
	pickUpNumber, _ := reader.ReadString('\n')
	wipeNumber := strings.TrimSpace(pickUpNumber)
	CaracterNumber := len(wipeNumber) - 1
	validaLetra, _ := strconv.Atoi(wipeNumber)
	for {
		if validaLetra > 0 || string(wipeNumber[5]) == "-" {
			if CaracterNumber < 7 {
				fmt.Println("Numero com caracteres insuficientes")
				break
			} else if CaracterNumber == 7 {
				wipeNumber = string(wipeNumber) + "3"
				fmt.Print("telefone:", pickUpNumber)
				fmt.Println("Telefone possui 7 dígitos. Vou acrescentar o digito três na frente.")
				fmt.Print("Telefone corrigido sem formatação:", pickUpNumber)
				fmt.Print("Telefone corrigido com formatação:", wipeNumber)
				break
			} else if CaracterNumber == 8 {
				fmt.Print("telefone:", pickUpNumber)
				break
			} else if CaracterNumber == 9 {
				if string(wipeNumber[5]) == "-" {
					fmt.Println("telefone:", pickUpNumber)
					break
				}
			}
		} else {
			fmt.Println("Numero invalido ")
			break
		}
	}
}
