package main

import "fmt"

func main() {

	// string
	// veja mais sobre strings em: https://blog.golang.org/strings
	name := "wilson"
	fmt.Println("Meu nome é ", name)

	// strings no fundo são arrays
	fmt.Printf("String como array: %v\n", "teste"[0])

	// sua estrutura é parecida com:
	type asString struct {
		length  uint64
		pointer [5]byte
	}

	// +----------------------------------------------------------------+----------------------------------------------------------------+
	// |                                                                |                                                                |
	// +------+---------------------------------------------------------+-------+--------------------------------------------------------+
	// 		  |                                                                 |
	// 		  |                                                                 |
	// 		  |                                                                 |
	// 		  |                                                                 +-------> +--------+
	// 		  |                                                                           |byte    |
	// 		  |                                                                           +--------+
	// 		  |
	// 		  |                                                                           .
	// 		  |                                                                           .
	// 		  |                                                                           .
	// 		  |
	// 		  |                                                                           +--------+
	// 		  |                                                                           |byte    |
	// 		  |                                                                           +--------+
	// 		  |
	// 		  |
	// 		  |
	// 		  +----------> +----------------------------------------------------------------+
	// 					   |uint64                                                          |
	// 					   +----------------------------------------------------------------+
	//
	// veja mais em: https://research.swtch.com/godata

	// lendo strings do STDIN
	fmt.Print("STDIN: ")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("Você escreveu: %s\n", input)

	// strings são imutáveis, as operações com elas sempre retornam uma nova string

}
