package main

import (
	"fmt"
	"math/cmplx"
)

// Nesse exemplo, vamos ver alguns tipos simples de variáveis em Go. Existem outros tipos além dos expostos
//    aqui, mas para entender como os outros tipos funcionam, é necessário conhecer esses.
// Basicamente, os tipos tratados aqui são:
// - Boleanos
// - Inteiros (que variam de 8 até 64 bits)
// - Vígula flutuante (e uma explicação sobre seu funcionamento)
// - Complexos

// Referências:
// 1. https://research.swtch.com/godata
// 2. http://perso.ens-lyon.fr/jean-michel.muller/goldberg.pdf
// 3. https://www.ardanlabs.com/blog/2013/08/gustavos-ieee-754-brain-teaser.html
// 4. https://ieeexplore.ieee.org/document/30711 (IEEE-754)
// 5. https://en.wikipedia.org/wiki/Unum_(number_format)

func main() {
	// booleanos
	var boleano bool
	fmt.Println("Valor boleano padrão:", boleano)
	verdade := true
	mentira := false
	fmt.Println("Verdade é igual a mentira:", verdade == mentira)

	// inteiros
	// Declarando uma variável utilizando o modelo pré-declarativo.
	var x int32
	fmt.Println("Valor int32 padrão:", x)
	x = 10
	fmt.Printf("O valor de 'x' é %d\n", x)

	// inteiros na memória
	// +--+
	// |10|
	// +--+
	// +-----------+    +--------------------------------+
	// |00 00 00 0A| =  |00000000000000000000000000001010|
	// +-----------+    +--------------------------------+
	//
	// saiba mais em [1]

	// Declarando outra variável utilizando o modelo auto-declarativo.
	//   Repare que o tipo padrão pode variar de int64 ou int32 dependendo da arquitetura.
	a := 20
	fmt.Printf("O valor de 'a' era %d\n", a)
	a = 30
	fmt.Printf("O valor de 'a' agora é %d\n", a)

	// vírgula flutuante
	// Os números em vírgula flutuante possuem uma representação aproximada, uma vez que não é possível representá-los de forma exata
	//   (com um tamanho fixo) em binário. Dessa forma, como esses números utilizam uma representação aproximada, seu valor armazenado
	//   na memória pode conter erros e esses podem ser acumulados de forma significativa durante as operações matemáticas.
	var decimo float32 = 0.1
	fmt.Printf("O valor do decimo é:\n%0.6f\n", decimo)
	// O mesmo número com 20 casas decimais mostra o erro acumulado.
	fmt.Printf("O valor do decimo é:\n%0.20f\n", decimo)

	// vírgula flutuante na memória
	//
	// A representação em vírgula flutuante é feita armazenando o número como se fosse em notação científica. Ou seja, teremos uma
	//   mantissa e um expoente, além do sinal de ambos. O número 0.1 seria representado da seguinte forma:
	//
	// 0.1? != 13421773p-27 = 13421773 * 2^-27
	//
	// Mas ao observamos como o número 0.1 está na memória, obtemos a seguinte informação:
	//   Para imprimir os bits de um float32, use a expressão a seguir:
	//   fmt.Printf("%032b\n", math.Float32bits(decimo))
	//
	// +-----------+    +----------------------------------+
	// |3d cc cc cd| =  |0 01111011 10011001100110011001101|
	// +-----------+    +----------------------------------+
	//                   0 123-127=-4
	//                              2^(-1) + 2^(-4) + 2^(-5) + 2^(-8) + 2^(-9) + 2^(-12) + 2^(-13) + 2^(-16) + 2^(-17) + 2^(-20) + 2^(-21) + 2^(-23)
	//                              = 0,6000000238
	//                   1,6000000238 * 2 ^(-4) = 0,1000000015
	//
	// Veja um passo a passo sobre como obter o número de sua representação em [3].
	// Veja mais sobre números em "virgula flutuante" em [2].

	// vírgula flutuante e as operações matemáticas
	// A aritmética muda uma vez que a representação dos números carrega erros.
	//   Repare que a ordem das operações altera o resultado!
	g := 1234.567
	r := 45.67834
	i := 0.0004
	s := (g + r) + i
	u := g + (r + i)
	fmt.Println("S é igual a U:", s == u)

	// float64
	// Nessa representação temos mais bits disponíveis, logo o erro que existe, é menor.
	pi := 3.14
	fmt.Printf("O valor de PI é %0.6f\n", pi)
	fmt.Printf("O valor de PI é %0.20f\n", pi)

	// Para além do vírgula flutuante existem formatos mais modernos. Caso se interesse pelo assunto,
	//   leia mais em [5].

	// complexos
	// São a representação dos números com uma parte real e outra imaginária. Como é formado de dois números
	//   (a parte real e a parte imaginária), pode ter tamanho de 64 ou 128 bits.
	// As operações matemáticas podem ser encontradas na biblioteca math/cmplx.
	h := complex(3, 4)
	fmt.Println("A hipotenusa de pitágoras:", cmplx.Abs(h))
	fmt.Println("A parte real é:", real(h))

	// Descobrindo tipos: para fazer isso, é possível utilizar o %T.
	fmt.Printf("verdade is of type %T\n", verdade)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("decimo is of type %T\n", decimo)
	fmt.Printf("pi is of type %T\n", pi)
	fmt.Printf("h is of type %T\n", h)
}
