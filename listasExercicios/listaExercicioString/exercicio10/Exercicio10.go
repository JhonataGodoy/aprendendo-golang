package main

//Escreva um programa que solicite ao usuário a digitação de um número até 99 e imprima-o na tela por extenso.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your Cpf:")
	pickUpCpf, _ := reader.ReadString('\n')
	wipeCpf := strings.TrimSpace(pickUpCpf)
	fmt.Println(wipeCpf)

}
