package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(Ola("mundo"))
}

func Ola(nome string) string {
	return "Olá, " + nome
}

func TestOla(t *testing.T) {
	resultado := Ola("Chris")
	esperado := "Olá, Chris"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
