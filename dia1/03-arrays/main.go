package main

import "fmt"

func main() {
	// arrays
	// Possuem tamanho fixo e são valores. Podem ser de todos os tipos.
	message := [5]byte{'h', 'e', 'l', 'l', 'o'}
	// +---------+
	// |h|e|l|l|o|
	// +---------+
	fmt.Printf("%v\n", message)

	primos := [6]uint32{1, 2, 3, 5, 7, 11}
	// +------------+
	// |1|2|3|5|7|11|
	// +------------+
	fmt.Printf("%v\n", primos)

}
