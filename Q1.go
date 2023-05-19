package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	matrix := Matrix{3, 3, [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}}

	fmt.Println("First Matrix")
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			fmt.Print(matrix.Elements[row][column], " ")
		}
		fmt.Print("\n")
	}

	matrix.GetRows()
	matrix.GetCol()

	fmt.Println("Changed first element of first matrix")
	matrix.setElement(1, 1, 10)
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			fmt.Print(matrix.Elements[row][column], " ")
		}
		fmt.Print("\n")
	}
	// matrix.printElements()
	fmt.Println("Second Matrix")
	matrix2 := Matrix{3, 3, [][]int{
		{11, 21, 31},
		{43, 53, 63},
		{73, 83, 39},
	}}
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			fmt.Print(matrix2.Elements[row][column], " ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\nAdded Matrix1 and Matrix2\n")
	final := matrix.Add(matrix2)
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			fmt.Print(final.Elements[row][column], " ")
		}
		fmt.Print("\n")
	}
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			fmt.Print(matrix.Elements[row][column], " ")
		}
		fmt.Print("\n")
	}

	EncodeJson(final)

}

type Matrix struct {
	Rows     int
	Columns  int
	Elements [][]int
}

// func Adder(matrix Matrix, matrix2 Matrix) {
// 	var m1 = matrix.Elements
// 	var m2 = matrix2.Elements

// 	var a [3][3]int
// 	for row := 0; row < 3; row++ {
// 		for column := 0; column < 3; column++ {
// 			a[row][column] = m1[row][column] + m2[row][column]
// 			fmt.Print(a[row][column], " ")
// 		}
// 		fmt.Print("\n")
// 	}

// }

func (matrix Matrix) Add(matrix2 Matrix) Matrix {
	// result := matrix
	var result Matrix

	for row := 0; row < matrix.Rows; row++ {
		var row1 []int
		for column := 0; column < matrix.Columns; column++ {
			row1 = append(row1, matrix.Elements[row][column]+matrix2.Elements[row][column])

		}
		result.Elements = append(result.Elements, row1)

	}
	return result

}

func (m Matrix) GetRows() {
	fmt.Println("Number of Rows: ", m.Rows)
}

func (m Matrix) GetCol() {
	fmt.Println("Number of Columns: ", m.Columns)
}

func (m *Matrix) setElement(i int, j int, finalVal int) {
	m.Elements[i-1][j-1] = finalVal
}

func EncodeJson(matrix Matrix) {
	finalJson, err := json.Marshal(matrix.Elements)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
