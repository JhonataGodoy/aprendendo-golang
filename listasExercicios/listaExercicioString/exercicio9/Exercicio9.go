package main

//Desenvolva um programa que solicite a digitação de um número de CPF no formato xxx.xxx.xxx-xx e
//indique se é um número válido ou inválido através da validação dos dígitos verificadores edos
//caracteres de formatação.

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
	fmt.Println("enter your Cpf:")
	pickUpCpf, _ := reader.ReadString('\n')
	wipeCpf := strings.TrimSpace(pickUpCpf)

	var pointValidation int
	if len(wipeCpf) == 14 {
		pointValidation += 1
	}
	if string(wipeCpf[3]) == "." {
		pointValidation += 1
	}
	if string(wipeCpf[7]) == "." {
		pointValidation += 1
	}
	if string(wipeCpf[11]) == "-" {
		pointValidation += 1
	}

	var validZero int
	for i := 0; i <= 13; i++ {
		caractersConvert := string(wipeCpf[i])
		intNumber, _ := strconv.Atoi(caractersConvert)

		if intNumber > 0 {
		} else if intNumber == 0 {
			validZero += 1
			if validZero > 3 {
				pointValidation = 0
			} else if validZero == 3 {
				pointValidation += 1
			}
		}
	}

	if pointValidation == 5 {
		fmt.Println("number válid")
	} else {
		fmt.Println("number invalid")
	}
}
