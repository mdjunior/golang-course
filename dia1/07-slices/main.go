package main

import "fmt"

func main() {
	// slices - são listas que podem crescer de tamanho
	//   lembre que arrays tem tamanho fixo
	linguagens := []string{"Perl", "python"}

	// slices não guardam valores e sim referencias para os arrays que os formam
	todasAsLinguagens := [6]string{
		"asm",
		"fortran",
		"C",
		"C++",
		"Ada",
		"Java",
	}
	civilUFRJ := todasAsLinguagens[1:3] // o slice civilUFRJ aponta para um pedaço do array todasAsLinguagens
	fmt.Printf("todasAsLinguagens: %v\n", todasAsLinguagens)
	fmt.Printf("civilUFRJ: %v\n", civilUFRJ)

	// como o slice aponta para o array, se alterarmos em um, alteramos no outro
	// alterando de C para Python
	todasAsLinguagens[2] = "python"
	fmt.Printf("todasAsLinguagens após alteração: %v\n", todasAsLinguagens)
	fmt.Printf("civilUFRJ após alteração: %v\n", civilUFRJ)

}
