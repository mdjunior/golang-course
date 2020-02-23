package main

import (
	"errors"
	"fmt"
)

// ErrNotAllowed é chamado quando o menor tem menos de 19 anos
var ErrNotAllowed = errors.New("Não permitida a entrada")

func entryInParty(age int) (string, error) {
	if age < 18 {
		return "", ErrNotAllowed
	}

	return "TICKET", nil
}

func main() {
	ticket, err := entryInParty(17)

	// Erros são valores, podemos verificar depois que chamamos a função
	if err != nil {
		fmt.Println("Falha ao entrar na Balada: ", err)
	}

	// É possível rastrear os erros no Go 1.13. Veja mais em: https://blog.golang.org/go1.13-errors

	ticket, err = entryInParty(17)
	if err == ErrNotAllowed {
		fmt.Println("Chame os pais da criança!")
	}

	fmt.Println("Ticket liberado", ticket)
}
