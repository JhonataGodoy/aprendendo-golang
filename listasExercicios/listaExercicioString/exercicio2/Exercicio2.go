package main

//Faça um programa que permita ao usuário digitar o seu nome e em seguida mostre o nome do usuário de trás
//para frente utilizando somente letras maiúsculas. Dica: lembre−se que ao informar o nome o usuário pode digitar
//letras maiúsculas ou minúsculas.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// finish
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("your name: ")
	pegaNome, _ := reader.ReadString('\n')
	stringNome := strings.TrimSpace(pegaNome)

	for i := len(stringNome); i > 0; i-- {
		fmt.Print(strings.ToUpper(string(stringNome[i-1])))
	}
}
