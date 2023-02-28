package main

//Desenvolva um jogo da forca. O programa terá uma lista de palavras lidas de um arquivo texto
//e escolherá uma aleatoriamente. O jogador poderá errar 6 vezes antes de ser enforcado.

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	abreviaCaminho := "/Users/jhgodoy/Documents/Transport Core/aprendendo-golang/listasExercicios/listaExercicioString/exercicio11/Exercicio11.txt"
	openFile, err := os.Open(abreviaCaminho)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())    // inicializa numeros aleatorios
	defer openFile.Close()              // fecha o arquivo lido na ultima linha da func
	reader := bufio.NewReader(openFile) // le o arquivo
	var pickUpWords []string
	for true {
		pickUpString, err := reader.ReadString('\n') // Aqui verifica se o arquivo existe, encontra a string vazia ' '
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		pickUpString = strings.Trim(pickUpString, "\n") // tira a string que eu quero, no caso o \n
		pickUpWords = append(pickUpWords, pickUpString) // soma
	}
	numberElements := len(pickUpWords)
	randomNumber := rand.Intn(numberElements) // numeros aleatorios

	var pickUpAllPhrase string
	var error int

	for i := len(pickUpWords[randomNumber]) - 1; i >= 0; i-- {
		var letraAcerto int
		readerD := bufio.NewReader(os.Stdin)
		fmt.Println("Digite uma letra:")
		pickUpLetra, _ := readerD.ReadString('\n')
		wipeLetra := strings.TrimSpace(pickUpLetra)

		pickUpAllPhrase = string(pickUpWords[randomNumber][i])
		if pickUpAllPhrase == wipeLetra {
			letraAcerto += 1
		}
		if letraAcerto == 0 {
			error += 1
		}
		fmt.Println("Quantidade de erros:", error, "vezes")
	}
	fmt.Println(pickUpWords[randomNumber])
}
