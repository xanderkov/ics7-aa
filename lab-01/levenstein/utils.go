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

