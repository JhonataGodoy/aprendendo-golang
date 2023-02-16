package main

//Desenvolva um jogo da forca. O programa terá uma lista de palavras lidas de um arquivo texto
//e escolherá uma aleatoriamente. O jogador poderá errar 6 vezes antes de ser enforcado.

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	abreviaCaminho := "/Users/jhgodoy/Documents/Transport Core/aprendendo-golang/listasExercicios/listaExercicioString/exercicio11/Exercicio11.txt"
	file, err := ioutil.ReadFile(abreviaCaminho)
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	fmt.Println(text)

}
