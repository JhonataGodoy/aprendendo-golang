package main

//Faça um programa que solicite o nome do usuário e imprima-o na vertical.

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

	for i := 0; i < len(stringNome); i++ {
		fmt.Println(string(stringNome[i]))
	}
}
