package main

// Modifique o programa anterior de forma a mostrar o nome em formato de escada.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("your name: ")
	pegaNome, _ := reader.ReadString('\n')
	stringNome := strings.TrimSpace(pegaNome)
	var escadaNome string
	for i := 0; i < len(stringNome); i++ {
		escadaNome += string(stringNome[i])
		fmt.Println(escadaNome)
	}
}
