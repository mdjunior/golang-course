package main

import "fmt"

func main() {

	// Vimos anteriormente que strings podem ser entendidas como uma lista (de tamanho fixo) de bytes, e que cada byte possui 8 bits.
	//   Em Go, todo o código fonte é escrito em UTF-8 e, dessa forma, é necessário assumir que um caractere em uma string pode ser codificado usando UTF-8.
	//   Dizemos que esse caractere "pode" assumir a codificação UTF-8 pois ele também pode não assumir, uma vez que é possível declarar uma string como
	//     uma sequência de bytes.
	//   Sendo uma string uma lista de bytes, devemos ter em mente que uma representação em UTF-8 por ter de 1 até 4 bytes, ou seja, podemos ter uma string
	//     que tenha uma lista de 4 bytes mas apenas um caractere.
	//   Para não confundir os conceitos, evitamos usar a palavra "caractere" e procuramos usar "code points" (denominação utilizada pelo Unicode).
	//   Em Go, o tipo que representa os "code points" são `rune`.

	// Referências:
	// - https://www.ime.usp.br/~pf/algoritmos/apend/unicode.html
	// - https://blog.golang.org/strings

	// Nesse exemplo podemos ver como apenas um code point (rune) utiliza 4 bytes.
	stringEspecial := "😅"
	for index := 0; index < len(stringEspecial); index++ {
		fmt.Printf("%d: %v\n", index, stringEspecial[index])
	}

	// Nesse exemplo podemos ver como ler a string com 1 rune por vez.
	stringRune := "😏😖😞"
	for index, char := range stringRune {
		fmt.Printf("%d: %v\n", index, char)
	}

	frase := `minha frase muito grande com muitas letras`
	countA := 0
	for _, char := range frase {
		if char == rune('a') {
			countA++
		}
	}
	fmt.Printf("Número de A's: %d\n", countA)

}
