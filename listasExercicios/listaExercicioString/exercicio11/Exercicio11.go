package main

//Desenvolva um jogo da forca. O programa terá uma lista de palavras lidas de um arquivo texto
//e escolherá uma aleatoriamente. O jogador poderá errar 6 vezes antes de ser enforcado.

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	pickUpWordTxt, _ := ioutil.ReadFile("Words.txt")
	pickUpSlice := strings.Fields(string(pickUpWordTxt))
	fmt.Println(pickUpSlice)

}
