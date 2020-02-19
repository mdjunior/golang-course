package main

import "fmt"

const pi = 3.14
const pii = "tres virgula quatorze"

// contantes são tipos "untyped", de forma que seu uso é mais flexível

func main() {
	fmt.Println("O valor de PI é", pi)
	fmt.Println("O valor de PI + complex(3, 0) é", pi+complex(3, 0))
	fmt.Println("O valor de PI + float32(3) é", pi+float32(3))

	// fazendo o mesmo com variáveis normais, é necessário declarar a conversão entre formatos
	// piFloat := 3.14
	// fmt.Println("O valor de PI (como float) + float32(3) é", piFloat+float64(3))
}
