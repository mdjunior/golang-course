package main

import "fmt"

func main() {
	// maps - hash table
	// maps são representados como referências, ou seja, se ele não estiver definido, pode ser nil
	people := map[string]int{
		"wilson": 25,
		"andre":  26,
	}

	// adiciona valor no map
	people["junior"] = 10

	// itera sobre tudo que está no map
	//   a ordem dos maps não é garantida
	for name, age := range people {
		fmt.Printf("Chave %s do mapa valor: %d\n", name, age)
	}

	// lendo entradas que não existem
	test := people["coelhinho da pascoa"]
	fmt.Printf("Coelho: %d\n", test)

	// apagando entrada
	delete(people, "wilson")
	fmt.Printf("Posição não alocada: %d\n", people["wilson"])

	// testando se posição existe
	test, ok := people["coelhinho da pascoa"]
	fmt.Printf("Coelho: %d, %v\n", test, ok)

	// saiba mais sobre maps em: https://blog.golang.org/go-maps-in-action
}
