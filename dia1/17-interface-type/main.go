package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// An interface type has an dinamic type inside

// Stringer is an interface with String() method
type Stringer interface {
	String() string
}

// ToString returns string for an interface
func ToString(any interface{}) string {
	// Altera o tipo da interface para ver se é Stringer
	if v, ok := any.(Stringer); ok {
		return v.String()
	}

	v := reflect.TypeOf(any).Kind()
	if v == reflect.Int {
		return strconv.Itoa(any.(int))
	}
	if v == reflect.Float64 {
		return strconv.FormatFloat(any.(float64), 'E', -1, 32)
	}

	return "???"
}

// Binary é um exemplo que implementa a interface Stringer
type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(uint64(i), 2)
}

func main() {
	var zerosUNs Binary = 100
	fmt.Println(zerosUNs.String())

	var zerosUNsInterface interface{} = 100
	fmt.Println(ToString(zerosUNsInterface))

	// Como Binary implementa o método String(), podemos transformar Binary em Stringer
	s := Stringer(zerosUNs)
	fmt.Printf("%v\n", s)

	// como a estrutura de uma interface parece
	// type iface struct {
	// 	tab  *itable  -> type e lista de funções associadas a interface
	// 	data *pointer -> ponteiro para os dados
	// }
	// você pode conhecer mais em: https://github.com/teh-cmc/go-internals/blob/master/chapter2_interfaces/README.md
	//   https://research.swtch.com/interfaces
	zerosUNs = 10
	fmt.Printf("%v\n", s)
}
