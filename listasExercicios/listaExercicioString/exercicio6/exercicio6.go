package main

//Faça um programa que solicite a data de nascimento (dd/mm/aaaa) do usuário e imprima a data com o nome do mês por extenso.

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// finish
func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("date of birth: (dd/mm/aaaa):")
	pickUpDate, _ := reader.ReadString('\n')
	wipeDate := strings.TrimSpace(pickUpDate)

	pickUpMonth := string(wipeDate[3:5])

	if pickUpMonth == "01" {
		pickUpMonth = "janeiro"
	} else if pickUpMonth == "02" {
		pickUpMonth = "fevereiro"
	} else if pickUpMonth == "03" {
		pickUpMonth = "março"
	} else if pickUpMonth == "04" {
		pickUpMonth = "abril"
	} else if pickUpMonth == "05" {
		pickUpMonth = "maio"
	} else if pickUpMonth == "06" {
		pickUpMonth = "junho"
	} else if pickUpMonth == "07" {
		pickUpMonth = "julho"
	} else if pickUpMonth == "08" {
		pickUpMonth = "agosto"
	} else if pickUpMonth == "09" {
		pickUpMonth = "setembro"
	} else if pickUpMonth == "10" {
		pickUpMonth = "outubro"
	} else if pickUpMonth == "11" {
		pickUpMonth = "novembro"
	} else if pickUpMonth == "12" {
		pickUpMonth = "dezembro"
	}

	fmt.Println("you were born in", string(wipeDate[0:2]), "from", pickUpMonth, "from", string(wipeDate[6:10]))

}
