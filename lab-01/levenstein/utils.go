package levenstein
import "fmt"

type MInt [][]int

func ReadWord() string {
	var word string
	fmt.Scanln(&word)

	return word
}

func makeMatrix(n, m int) MInt {
	matrix := make(MInt, n + 1)
	
	for i := range matrix {
		matrix[i] = make([]int, m + 1)
	}

	for i := 0; i < m + 1; i++ {
		matrix[0][i] = i
	}

	for i := 0; i < n + 1; i++ {
		matrix[i][0] = i
	}
	return matrix
}


func makeMatrixInf(n, m int) MInt {
	matrix := make(MInt, n + 1)
	
	for i := range matrix {
		matrix[i] = make([]int, m + 1)
	}

	for i := 0; i < n + 1; i++ {
		for j := 0; j < m + 1; j++ {
			matrix[i][j] = -1
		}
	}
	return matrix
}

func (mat MInt) PrintMatrix() {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			fmt.Printf("%3d ", mat[i][j])
		}
		fmt.Printf("\n")
	}
}

func getMinOfValues(values ...int) int {
	min := values[0]

	for _, i := range values {
		if min > i {
			min = i
		}
	}

	return min
}

func getMaxOf2Values(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}