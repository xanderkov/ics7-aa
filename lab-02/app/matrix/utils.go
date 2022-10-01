package matrix
import "fmt"
import "math/rand"

type MInt struct {
	mat  [][]int
	n, m int
}


func (mat MInt) PrintMatrix() {
	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m; j++ {
			fmt.Printf("%3d ", mat.mat[i][j])
		}
		fmt.Printf("\n")
	}
}

func ReadMatrix(n, m int) MInt {
	mat := formResMat(n, m)

	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m; j++ {
			fmt.Scanf("%d", &mat.mat[i][j])
		}
	}

	return mat
}

func ReadNum() int {
	var num int

	fmt.Scanln(&num)

	return num
}

func randomFill(mat MInt, max int) {
	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m; j++ {
			mat.mat[i][j] = rand.Intn(max)
		}
	}
}

func formResMat(n, m int) MInt {
	var rmat MInt

	rmat.n = n
	rmat.m = m
	rmat.mat = make([][]int, rmat.n)

	for i := range rmat.mat {
		rmat.mat[i] = make([]int, rmat.m)
	}

	return rmat
}
