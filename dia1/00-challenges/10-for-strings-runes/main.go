package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Você precisa fazer um programa que calcule por quantos vales um sistema de posicionamento passou.
// Para registrar uma inclinação nagativa, o sistema utiliza a letra D (down) e para registrar uma inclinação
// positiva, ele utiliza a letra U (up).
// A seguinte sequência mostra apenas 1 vale:
// UUUUDDDDUDUDDDUU

// Podemos separar ela da seuinte forma:
// UUUUDDDD_UD_UD____________DDUU
//  |        |                 |
// montanha  outras montanha   |
// 						   aqui temos um vale!

// Referência: https://www.hackerrank.com/challenges/counting-valleys/problem

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	var valleys int32

	return valleys
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	result := countingValleys(n, s)

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
