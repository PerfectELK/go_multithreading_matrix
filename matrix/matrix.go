package matrix

import (
	"errors"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type row struct {
	rowNumber int
	rowArr    []float64
}

func Multiply(
	matrix1 [][]float64,
	matrix2 [][]float64,
	threadNumber ...int,
) ([][]float64, error) {
	e := checkMatrix(matrix1, matrix2)
	if e != nil {
		return nil, e
	}

	threads := 2
	if len(threadNumber) > 0 {
		threads = threadNumber[0]
	}

	calculated := make(map[int]bool)
	var wg sync.WaitGroup
	var mu sync.Mutex

	var rowsArr []*[]row
	for i := 0; i < threads; i++ {
		wg.Add(1)
		var rows []row
		rowsArr = append(rowsArr, &rows)
		go calcRow(
			&wg,
			&mu,
			&calculated,
			&matrix1,
			&matrix2,
			&rows,
		)
	}
	wg.Wait()

	var newMatrixRows []row

	for _, arr := range rowsArr {
		for _, r := range *arr {
			newMatrixRows = append(newMatrixRows, r)
		}
	}

	sort.SliceStable(newMatrixRows, func(i, j int) bool {
		return newMatrixRows[i].rowNumber < newMatrixRows[j].rowNumber
	})

	var newMatrix [][]float64

	for _, r := range newMatrixRows {
		newMatrix = append(newMatrix, r.rowArr)
	}

	return newMatrix, nil
}

func calcRow(
	wg *sync.WaitGroup,
	mu *sync.Mutex,
	calculated *map[int]bool,
	matrix1 *[][]float64,
	matrix2 *[][]float64,
	rows *[]row,
) {
	defer wg.Done()
	for i := 0; i < len(*matrix1); i++ {
		mu.Lock()
		_, ok := (*calculated)[i]
		if ok {
			mu.Unlock()
			continue
		}
		matrix1Row := (*matrix1)[i]
		(*calculated)[i] = true
		mu.Unlock()

		var r row
		r.rowNumber = i
		var rowArr []float64
		for k := 0; k < len((*matrix2)[0]); k++ {
			var cell float64
			for j := 0; j < len(matrix1Row); j++ {
				matrix1RowValue := matrix1Row[j]
				cell += matrix1RowValue * (*matrix2)[j][k]
			}

			rowArr = append(rowArr, cell)
		}
		r.rowArr = rowArr
		*rows = append(*rows, r)
	}
}

func checkMatrix(matrix1 [][]float64, matrix2 [][]float64) error {
	var colNumber int = 0
	var rowNumber int = 0

	for _, row := range matrix1 {
		if colNumber == 0 {
			colNumber = len(row)
		}
		if len(row) != colNumber {
			return errors.New("matrix 1 is wrong")
		}
		colNumber = len(row)
	}
	rowNumber = len(matrix2)
	if rowNumber != colNumber {
		return errors.New("number of cols matrix1 is not equivalent to number of rows matrix2")
	}

	var colNumber2 int = 0
	for _, row := range matrix2 {
		if colNumber2 == 0 {
			colNumber2 = len(row)
		}
		if len(row) != colNumber2 {
			return errors.New("matrix 2 is wrong")
		}
		colNumber2 = len(row)
	}

	return nil
}

func Generate(rows int, cols int) [][]float64 {
	var matrix [][]float64

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < rows; i++ {
		matrix = append(matrix, []float64{})
		for j := 0; j < cols; j++ {
			matrix[i] = append(matrix[i], rand.NormFloat64()*100)
		}
	}

	return matrix
}
