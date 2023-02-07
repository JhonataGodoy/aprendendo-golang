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
	escadaNome := string(stringNome)
	tiraCaracter := len(stringNome)

	for i := 0; i < len(stringNome); i++ {
		fmt.Println(string(escadaNome[0:tiraCaracter]))
		if len(stringNome) >= 0 {
			tiraCaracter = len(stringNome) - i
		}
	}
}
