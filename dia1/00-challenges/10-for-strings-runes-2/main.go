package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Você precisa fazer um programa para renderizar um retangulo quando uma palavra é selecionada no editor de texto.
// Para renderizar esse retângulo, você deve calcular a área utilizada e para isso, você possui a sua disposição
//   o tamanho de cada letra do alfabeto (no slice h) e a palavra em questão (na string word).
//
// Como exemplo, para o slice h a seguir, a palavra abc terá uma área de 9 unidades quadradas (e portanto, a função
// designer deve retorna 9).
// h = [ 1 3 1 3 1 4 1 3 2 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 ]
// word = "abc"
//
// Resultado esperado: 9

// Referência: https://www.hackerrank.com/challenges/designer-pdf-viewer/problem

// Complete the designerPdfViewer function below.
func designerPdfViewer(h []int32, word string) int32 {
	var area int32

	return area
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	hTemp := strings.Split(readLine(reader), " ")

	var h []int32

	for i := 0; i < 26; i++ {
		hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
		checkError(err)
		hItem := int32(hItemTemp)
		h = append(h, hItem)
	}

	word := readLine(reader)

	result := designerPdfViewer(h, word)

	fmt.Printf("%d\n", result)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
