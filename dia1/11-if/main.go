package main

import (
	"fmt"
)

func main() {
	// if bool
	x := true
	if x == true {
		fmt.Println("x é true")
	}

	// if float
	y := 0.1
	z := 0.3 - 0.2
	if z == y {
		fmt.Println("O ponto flutuante deu certo")
	} else {
		fmt.Println("O ponto flutuante deu ruim")
	}

	// if string
	name := "manoel"
	surname := []byte{'m', 'a', 'n', 'o', 'e', 'l'}

	if string(surname) == name {
		fmt.Println("O nome é igual")
	}

}
