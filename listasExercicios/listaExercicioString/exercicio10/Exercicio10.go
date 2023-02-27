package main

//Escreva um programa que solicite ao usuário a digitação de um número até 99 e imprima-o na tela por extenso.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("number up to 99:")
	pickUpNumber, _ := reader.ReadString('\n')
	wipeNumber := strings.TrimSpace(pickUpNumber)

	if wipeNumber == "1" {
		fmt.Println("um")
	} else if wipeNumber == "2" {
		fmt.Println("dois")
	} else if wipeNumber == "3" {
		fmt.Println("tres")
	} else if wipeNumber == "4" {
		fmt.Println("quatro")
	} else if wipeNumber == "5" {
		fmt.Println("cinco")
	} else if wipeNumber == "6" {
		fmt.Println("seis")
	} else if wipeNumber == "7" {
		fmt.Println("sete")
	} else if wipeNumber == "8" {
		fmt.Println("oito")
	} else if wipeNumber == "9" {
		fmt.Println("nove")
	}

	if string(wipeNumber[0]) == "1" {
		if wipeNumber == "10" {
			fmt.Println("dez")
		} else if wipeNumber == "11" {
			fmt.Println("onze")
		} else if wipeNumber == "12" {
			fmt.Println("doze")
		} else if wipeNumber == "13" {
			fmt.Println("treze")
		} else if wipeNumber == "14" {
			fmt.Println("quatorze")
		} else if wipeNumber == "15" {
			fmt.Println("quinze")
		} else if wipeNumber == "16" {
			fmt.Println("dezesseis")
		} else if wipeNumber == "17" {
			fmt.Println("dezessete")
		} else if wipeNumber == "18" {
			fmt.Println("dezoito")
		} else if wipeNumber == "19" {
			fmt.Println("dezenove")
		}
	}

	var ExtensiveStart string
	if len(wipeNumber) == 2 {
		getsDecimal := string(wipeNumber[0])
		getsUnit := string(wipeNumber[1])
		convertDecimal, _ := strconv.Atoi(getsDecimal)
		convertUnit, _ := strconv.Atoi(getsUnit)
		if convertDecimal == 2 && string(wipeNumber[0]) != "1" {
			ExtensiveStart = "vinte"
		} else if convertDecimal == 3 {
			ExtensiveStart = "trinta"
		} else if convertDecimal == 4 {
			ExtensiveStart = "quarenta"
		} else if convertDecimal == 5 {
			ExtensiveStart = "cinquenta"
		} else if convertDecimal == 6 {
			ExtensiveStart = "sessenta"
		} else if convertDecimal == 7 {
			ExtensiveStart = "setenta"
		} else if convertDecimal == 8 {
			ExtensiveStart = "oitenta"
		} else if convertDecimal == 9 {
			ExtensiveStart = "noventa"
		}

		if string(wipeNumber[0]) != "1" {
			if convertUnit == 1 {
				ExtensiveStart += " e um"
			} else if convertUnit == 2 {
				ExtensiveStart += " e dois"
			} else if convertUnit == 3 {
				ExtensiveStart += " e tres"
			} else if convertUnit == 4 {
				ExtensiveStart += " e quatro"
			} else if convertUnit == 5 {
				ExtensiveStart += " e cinco"
			} else if convertUnit == 6 {
				ExtensiveStart += " e seis"
			} else if convertUnit == 7 {
				ExtensiveStart += " e sete"
			} else if convertUnit == 8 {
				ExtensiveStart += " e oito"
			} else if convertUnit == 9 {
				ExtensiveStart += " e nove"
			}
		}

	}

	fmt.Println(ExtensiveStart)

}
