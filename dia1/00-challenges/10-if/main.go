package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers.
//   Then print the respective minimum and maximum values as a single line of two space-separated long integers.
//   Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	var sum int64
	// Max and min needs be initialized
	var max int64 = int64(arr[0])
	var min int64 = int64(arr[0])

	for _, num := range arr {
		sum = sum + int64(num)
		if int64(num) > max {
			max = int64(num)
		}
		if int64(num) < min {
			min = int64(num)
		}
	}
	fmt.Printf("%d %d", sum-max, sum-min)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
