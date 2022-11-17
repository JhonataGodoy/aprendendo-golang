package inteiros

import (
	"testing"
)

func TestSomarFloat(t *testing.T) {
	var soma float64 = Adiciona(1.5, 2.7)
	esperado := 4.2

	if soma != esperado {
		t.Errorf("esperado '%f', resultado '%f'", esperado, soma)
	}
}

func TestSomarPositivoENegativo(t *testing.T) {
	soma := Adiciona(-2, 2)
	esperado := 0.0

	if soma != esperado {
		t.Errorf("esperado '%f', resultado '%f'", esperado, soma)
	}
}

func TestAdicionador(t *testing.T) {
	soma := Adiciona(4, 5)
	esperado := 9.0

	if soma != esperado {
		t.Errorf("esperado '%f', resultado '%f'", esperado, soma)
	}
}
