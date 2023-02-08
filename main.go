package main

import (
	"fmt"
	"main/matrix"
	"time"
)

func main() {

	//matrix1 := [][]float64{
	//	{1, 2, 2},
	//	{3, 1, 1},
	//}
	//
	//matrix2 := [][]float64{
	//	{4, 2},
	//	{3, 1},
	//	{1, 5},
	//}

	timeBeforeGenerate := float64(time.Now().UnixMilli()) / 1000
	matrix1 := matrix.Generate(1000, 1000)
	matrix2 := matrix.Generate(1000, 1000)
	timeAfterGenerate := float64(time.Now().UnixMilli()) / 1000

	_, e := matrix.Multiply(matrix1, matrix2, 10)
	timeAfterMultiply := float64(time.Now().UnixMilli()) / 1000

	fmt.Println("Between generate: ", timeAfterGenerate-timeBeforeGenerate, " sec")
	fmt.Println("Multiply time: ", timeAfterMultiply-timeAfterGenerate, " sec")

	if e != nil {
		fmt.Println(e)
	}

}
