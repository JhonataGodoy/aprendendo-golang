package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um Programa que peça as 4 notas bimestrais e mostre a média.

func main() {
	somaMedia := 0
	var notasB [4]int
	for i := 0; i < 4; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(" Digite a nota do", i+1, "bimestre: ")
		pegaNota, _ := reader.ReadString('\n')
		limpaNota := strings.TrimSpace(pegaNota)
		nota, err := strconv.Atoi(limpaNota)
		if err != nil {
			fmt.Println("O que você digitou não é um numero ")
			continue
		} else if nota < 0 {
			continue
		}
		notasB[i] = nota
		somaMedia = somaMedia + notasB[i]
	}

	fmt.Println("Media das notas bimestral:", somaMedia/4)

}
