package main

import "fmt"
import matrix "app/matrix"

func main() {
	fmt.Println("Умножение матриц\n")

	fmt.Println("Матриа А")
	fmt.Print("Введите количество строк: ")
	an := matrix.ReadNumber()
	fmt.Print("Введите количество столбцов: ")
	am := matrix.ReadNumber()
	fmt.Println("Введите матрицу: ")
	amat := matrix.ReadMatrix(an, am)

	fmt.Println("\n Матрица B")
	fmt.Print("Введите количество строк: ")
	bn := matrix.ReadNumber()

	if bn != am {
		fmt.Println("Количество строк матриц не совпадает")
		return
	}
	fmt.Print("Введите количество столбцов: ")
	bm := matrix.ReadNumber()
	fmt.Println("Введите матрицу: ")
	bmat := matrix.ReadMatrix(bn, bm)	

	fmt.Println("\n Резулитирующая матрица")
	fmt.Println("Обычное умножение: ")
	smmat := matrix.SimpleMult(amat, bmat)
	smmat.PrintMatirx()

	fmt.Println("\n Умножение по Винограду")
	wmmat := matrix.WinogradMult(amat, bmat)
	wmmat.PrintMatirx()

	fmt.Println("Умножение по Винограду (улучшенное): ")
	wimmat := matrix.WinogradMultBetter(amat, bmat)
	wimmat.PrintMatirx()
}