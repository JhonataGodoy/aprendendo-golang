package main

//Um palíndromo é uma seqüência de caracteres cuja leitura é idêntica se feita da direita para esquerda
//ou vice−versa. Por exemplo: OSSO e OVO são palíndromos. Em textos mais complexos os espaços e pontuação
//são ignorados. A frase SUBI NO ONIBUS é o exemplo de uma frase palíndroma onde os espaços foram ignorados.
//Faça um programa que leia uma seqüência de caracteres, mostre−a e diga se é um palíndromo ou não.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your Cpf:")
	pickUpPalim, _ := reader.ReadString('\n')
	wipePalim := strings.TrimSpace(pickUpPalim)

	leftRight := ""
	rightLeft := ""
	fmt.Println(wipePalim, leftRight, rightLeft)
}
