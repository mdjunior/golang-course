package main

import "fmt"

// Estruturas de dados permitem combinar os tipos pré-definidos existentes (built-in) de
//   forma a criar novos tipos mais adequados às aplicações finais.
// Entender as estruturas de dados nesse momento é importante pois existem outros tipos built-in que
//   são estruturas internas da linguagem (e existem como tipos para facilitar o uso).

// person é uma estrutura privada (ou seja, somente pode ser utilizada nesse package) que possui
//   os dados de uma pessoa.
type person struct {
	ID     uint64
	Age    uint32
	Active bool
}

// Student é uma estrutura pública (ou seja, que pode ser utilizada por outros pacotes, desde que importem esse).
//   Para caracterizar como uma estrutura pública, basta que a inicial seja maiúscula.
type Student struct {
	ID     uint64
	RG     person
	Active bool
}

func main() {
	// structs
	// Declaramos a variável person1 sem definir os valores iniciais.
	var person1 person
	fmt.Println(person1)

	// Declaramos a variável p fornecendo os nomes dos atributos.
	p := person{
		ID:     1,
		Age:    24,
		Active: true,
	}
	// Para acessar cada atributo, usamos `.`.
	fmt.Println("ID: ", p.ID)
	fmt.Println("Age: ", p.Age)
	fmt.Println("Active:", p.Active)

	// Declaramos a variável p2 passando apenas os valores na ordem correta. Essa é uma forma menos legível.
	p2 := person{2, 28, true}
	fmt.Println("ID: ", p2.ID)
	fmt.Println("Age: ", p2.Age)
	fmt.Println("Active:", p2.Active)

	// A estrutura aloca regiões contínuas de memória conforme for especificado. No caso de person, seria:
	// +-------------------------------------------------------------------------------------------------+-+
	// |uint64                                                          |uint32                          | | bool
	// +-------------------------------------------------------------------------------------------------+-+

	// structs and pointers
	// Podemos criar ponteiros para variáveis, permitindo que elas sejam passadas como referências a funções.
	//   Para criar um ponteiro, utilize o `&` na frente do nome da variável. Para acessar os atributos, basta usar o `.`:
	pointerToP := &p
	p.Age++
	fmt.Println("Pointer ID: ", pointerToP.ID)
	fmt.Println("Pointer Age: ", pointerToP.Age)
	fmt.Println("Pointer Active:", pointerToP.Active)

	// O ponteiro na memória apontará diretamente para a posição onde os dados estarão.
	// +----------------------------------------------------------------+
	// |                                                                | *person
	// +------------------------+---------------------------------------+
	// 							|
	// 							|
	// 							+---------> +-------------------------------------------------------------------------------------------------+-+
	// 										|uint64                                                          |uint32                          | | bool (person)
	// 										+-------------------------------------------------------------------------------------------------+-+

	// deep structs
	type classroom struct {
		teacher  person
		students [10]person
	}
	// Na memória, a estrutura classroom ficará da seguinte forma:
	// +---------------------------------------------------------------------------------------------------+
	// |uint64                                                          |uint32                          | | bool  (teacher)
	// +---------------------------------------------------------------------------------------------------+
	//
	// +---------------------------------------------------------------------------------------------------+
	// |uint64                                                          |uint32                          | | bool  (student 0)
	// +---------------------------------------------------------------------------------------------------+
	//
	//  .
	//  .
	//  .
	//
	// +---------------------------------------------------------------------------------------------------+
	// |uint64                                                          |uint32                          | | bool  (student 9)
	// +---------------------------------------------------------------------------------------------------+
	//

	// Podemos criar a variável goCourse do tipo classroom conforme já vimos antes:
	goCourse := classroom{
		p,
		[10]person{p, p2},
	}
	fmt.Printf("Classe: %v\n", goCourse)
	// Repare que nessa estrutura temos valores.

	// structs with pointers
	// Também podemos ter estruturas que usam ponteiros. Na estrutura do exemplo a seguir,
	//   ao invés de termos uma variável do tipo person e uma lista, temos dois ponteiros.
	type refClass struct {
		teacher  *person
		students *[10]person
	}

	// Na memória, essa estrutura fica organizada da seguinte forma:
	// +----------------------------------------------------------------+----------------------------------------------------------------+
	// |                                                                |                                                                |
	// +---------+------------------------------------------------------+----------+-----------------------------------------------------+
	//           |                                                                 |
	//           |                                                                 |
	//           |                                                                 |
	//           |                                                                 +-------> +-------------------------------------------------------------------------------------------------+-+
	//           |                                                               (student 0) |uint64                                                          |uint32                          | | bool
	//           |                                                                           +-------------------------------------------------------------------------------------------------+-+
	//           |
	//           |                                                                           .
	//           |                                                                           .
	//           |                                                                           .
	//           |
	//           |                                                                           +-------------------------------------------------------------------------------------------------+-+
	//           |                                                               (student 9) |uint64                                                          |uint32                          | | bool
	//           |                                                                           +-------------------------------------------------------------------------------------------------+-+
	//           |
	//           |
	//           |
	//           +----------> +-------------------------------------------------------------------------------------------------+-+
	//                        |uint64                                                          |uint32                          | | bool (teacher)
	//                        +-------------------------------------------------------------------------------------------------+-+

}
