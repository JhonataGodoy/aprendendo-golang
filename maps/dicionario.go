package main

import "errors"

type Dicionario map[string]string

var ErrNaoEncontrado = errors.New("não foi possível encontrar a palavra procurada")

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, ok := d[palavra]
	if !ok {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}
