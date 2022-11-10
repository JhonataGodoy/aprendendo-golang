package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Ola("jhonete", "francês")
		esperado := "Bonjour, jhonete"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em portugues", func(t *testing.T) {
		resultado := Ola("Chris", "portugues")
		esperado := "Olá, Chris"
		verificaMensagemCorreta(t, resultado, esperado)
	})

}
