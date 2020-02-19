package main

import "fmt"

func main() {
	// slices são estruturas de dados, como strings, mas podem crescer
	type sliceHeader struct {
		Length        int
		Capacity      int
		ZerothElement *byte
	}
	// quando criamos um slice, podemos definir seu tamanho e capacidade inicial
	primeNumbers := make([]int, 5, 10)
	fmt.Printf("len: %d, cap: %d\n", len(primeNumbers), cap(primeNumbers))

	primeNumbers[0] = 1
	primeNumbers[1] = 2
	primeNumbers[2] = 3
	primeNumbers[3] = 5
	primeNumbers[4] = 7
	fmt.Printf("len: %d, cap: %d, values: %v\n", len(primeNumbers), cap(primeNumbers), primeNumbers)

	primeNumbers = append(primeNumbers, 11)
	primeNumbers = append(primeNumbers, 13)
	fmt.Printf("len: %d, cap: %d, values: %v\n", len(primeNumbers), cap(primeNumbers), primeNumbers)

	primeNumbers = append(primeNumbers, 17)
	primeNumbers = append(primeNumbers, 19)
	primeNumbers = append(primeNumbers, 23)
	primeNumbers = append(primeNumbers, 29)
	fmt.Printf("len: %d, cap: %d, values: %v\n", len(primeNumbers), cap(primeNumbers), primeNumbers)

	// veja mais sobre slices em: https://blog.golang.org/slices
}
