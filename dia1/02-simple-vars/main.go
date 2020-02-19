package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// bool
	verdade := true
	mentira := false
	fmt.Println("Verdade é igual a mentira:", verdade == mentira)

	// int
	// declarando uma variável utilizando o modelo pré declarativo
	var x int // ou var x int = 10
	x = 10
	fmt.Printf("O valor de 'x' é %d\n", x)

	// +--+
	// |10|
	// +--+
	// +-----------+    +--------------------------------+
	// |00 00 00 0A| =  |00000000000000000000000000001010|
	// +-----------+    +--------------------------------+
	//
	// saiba mais em https://research.swtch.com/godata

	// declarando uma variável utilizando o modelo auto-declarativo
	a := 20
	fmt.Printf("O valor de 'a' era %d\n", a)
	a = 30
	fmt.Printf("O valor de 'a' agora é %d\n", a)

	// float32
	// veja mais sobre números em "virgula flutuante" em http://perso.ens-lyon.fr/jean-michel.muller/goldberg.pdf
	var decimo float32 = 0.1
	fmt.Printf("O valor do decimo é:\n%0.6f\n", decimo) // altere para 20 casas decimais

	// IEEE-754
	// 0.1? != 13421773p-27 = 13421773 * 2^-27
	// +-----------+    +----------------------------------+
	// |3d cc cc cd| =  |0 01111011 10011001100110011001101|
	// +-----------+    +----------------------------------+
	//                   0 123-127=-4
	//                              2^(-1) + 2^(-4) + 2^(-5) + 2^(-8) + 2^(-9) + 2^(-12) + 2^(-13) + 2^(-16) + 2^(-17) + 2^(-20) + 2^(-21) + 2^(-23)
	//                              = 0,6000000238
	//                   1,6000000238 * 2 ^(-4) = 0,1000000015
	//
	// veja um passo a passo sobre como obter o número de sua representação em https://www.ardanlabs.com/blog/2013/08/gustavos-ieee-754-brain-teaser.html
	//
	// Para imprimir a representação de um float, use a expressão a seguir:
	// fmt.Printf("%b\n", decimo)
	// Para imprimir os bits de um float32, use a expressão a seguir:
	// fmt.Printf("%032b\n", math.Float32bits(decimo))

	// a aritmética muda, repare que a ordem das operações altera o resultado!
	g := 1234.567
	r := 45.67834
	i := 0.0004
	s := (g + r) + i
	u := g + (r + i)
	fmt.Println("S é igual a U:", s == u)

	// float64
	pi := 3.14
	fmt.Printf("O valor de PI é %0.6f\n", pi)

	// complex
	h := complex(3, 4)
	fmt.Println("A hipotenusa de pitágoras:", cmplx.Abs(h))
	fmt.Println("A parte real é:", real(h))

	// descobrindo tipos
	fmt.Printf("verdade is of type %T\n", verdade)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("decimo is of type %T\n", decimo)
	fmt.Printf("pi is of type %T\n", pi)
	fmt.Printf("h is of type %T\n", h)
}
