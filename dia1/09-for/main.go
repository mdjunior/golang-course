package main

import "fmt"

func main() {
	// for - estrutura de repetição
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// for sem init e post
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// for que nem while
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	// for em slices
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for index, value := range pow {
		fmt.Printf("2^%d = %d\n", index, value)
	}

	// Loops em assembly: https://medium.com/a-journey-with-go/go-how-are-loops-translated-to-assembly-835b985309b3
}
