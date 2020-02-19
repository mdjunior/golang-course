package main

import "fmt"

// exemplo de função privada
func add(x uint32, y uint32) uint32 {
	return x + y // e se a soma der erro (se for maior que uint32?)
}

func main() {
	fmt.Println("O resultado de 10 + 20 é", add(10, 20))
}
