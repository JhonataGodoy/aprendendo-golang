package main

import "testing"

func Ola() string {
	return "Olá, mundo 2"
}

func TestOla(t *testing.T) {
	resultado := Ola()
	esperado := "Olá, mundo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
