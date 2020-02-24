package main

import "fmt"

func main() {
	// arrays
	// Possuem tamanho fixo e são valores (ou seja, quando passados para funções, todo seu valor é copiado).
	//   Podem ser de todos os tipos.

	// array de bytes
	message := [5]byte{'h', 'e', 'l', 'l', 'o'}
	// +---------+
	// |h|e|l|l|o|
	// +---------+
	fmt.Printf("%v\n", message)

	// array de uint32
	primos := [6]uint32{1, 2, 3, 5, 7, 11}
	// +------------+
	// |1|2|3|5|7|11|
	// +------------+
	fmt.Printf("%v\n", primos)

	// Cópia de message com posterior alteração. Lembre disso quando vermos slices!
	messageCopy := message
	messageCopy[0] = 'H'
	fmt.Printf("%v\n", message)
	fmt.Printf("%v\n", messageCopy)

	// array de tamanho automático
	semTamanho := [...]uint64{
		1,
		1,
		2,
		3,
		5,
		8,
		13,
		21,
	}
	fmt.Printf("Tamanho do array: %d\n", len(semTamanho))
}
