package main

import (
	"fmt"
	"math"
)

// exemplo de função retornando 2 valores
func add(x uint32, y uint32) (uint32, bool) {
	var sum uint64
	sum = uint64(x) + uint64(y)

	var err bool

	if sum > math.MaxUint32 {
		err = true
	}
	return uint32(sum), err // vamos ver erros mais para a frente
}

func main() {
	sum, err := add(1000, 20)
	if err == false {
		fmt.Println("O resultado de 10 + 20 é", sum)
	}
}
