package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Compare reviews, where 1st line is Alice and 2nd line is Bob. Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
	result := []int32{0, 0}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	var a []int32
	for i := 0; i < len(aTemp); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	bTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	var b []int32
	for i := 0; i < len(bTemp); i++ {
		bItemTemp, err := strconv.ParseInt(bTemp[i], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	result := compareTriplets(a, b)

	for i, resultItem := range result {
		fmt.Printf("%d", resultItem)

		if i != len(result)-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
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
