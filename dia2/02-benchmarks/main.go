package main

import "fmt"

const quantidadeRepeticoes = 5

// Repetir repete uma string um número constante de vezes
func Repetir(caractere string) string {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}

func main() {
	fmt.Println(Repetir("A"))
}
