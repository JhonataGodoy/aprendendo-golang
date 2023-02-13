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

// finish
func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your phone number:")
	pickUpNumber, _ := reader.ReadString('\n')
	wipeNumber := strings.TrimSpace(pickUpNumber)

	pointValidation := 0
	if len(wipeNumber) == 14 {
		pointValidation += 1
		//	confere se tem 14 caracteres ok
	}
	if string(wipeNumber[3]) == "." {
		pointValidation += 1
		//confere primeiro ponto ok
	}
	if string(wipeNumber[7]) == "." {
		pointValidation += 1
		//confere segundo ponto ok
	}
	if string(wipeNumber[11]) == "-" {
		pointValidation += 1
		//confere terceiro ponto ok
	}

	validZero := 0
	for i := 0; i <= 13; i++ {
		caractersConvert := string(wipeNumber[i])
		intNumber, _ := strconv.Atoi(caractersConvert)

		if intNumber > 0 {
			//valido, numero ok
		} else if intNumber == 0 {
			validZero += 1
			if validZero > 3 {
				pointValidation = 0
				// se tiver mais de 3 0 ele reinicia a variavel fazendo nao passar na validacao
			} else if validZero == 3 {
				pointValidation += 1
				// só pode ter 3 zeros entao se == 3 e passar ok
			}
		}
	}

	if pointValidation == 5 {
		fmt.Println("number válid")
	} else {
		fmt.Println("number invalid")
	}
}
