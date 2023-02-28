package main

//Dado uma string com uma frase informada pelo usuário (incluindo espaços em branco), conte:
//quantos espaços em branco existem na frase.
//quantas vezes aparecem as vogais a, e, i, o, u.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("say a phrase:")
	pickUpPhrase, _ := reader.ReadString('\n')
	convertMaius := strings.ToUpper(pickUpPhrase)

	var spaceA, spaceE, spaceI, spaceO, spaceU, whiteSpace int
	var pickUpAllPhrase string

	for i, _ := range pickUpPhrase {

		pickUpAllPhrase = string(convertMaius[i])
		if pickUpAllPhrase == "" {
			whiteSpace += 1
		} else if pickUpAllPhrase == "A" {
			spaceA += 1
		} else if pickUpAllPhrase == "E" {
			spaceE += 1
		} else if pickUpAllPhrase == "I" {
			spaceI += 1
		} else if pickUpAllPhrase == "O" {
			spaceO += 1
		} else if pickUpAllPhrase == "U" {
			spaceU += 1
		}

	}
	fmt.Println("space white:", whiteSpace)
	fmt.Println("space A:", spaceA, "space E:", spaceE, "spaceI:", spaceI, "space O:", spaceO, "space U:", spaceU)
}
