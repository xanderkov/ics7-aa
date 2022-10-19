package main

import (
	"fmt"
	"./matrix"
)

func main() {
	fmt.Println("Умножение матриц\n")

	fmt.Println("Матриа А")
	fmt.Print("Введите количество строк: ")
	an := matrix.ReadNum()
	fmt.Print("Введите количество столбцов: ")
	am := matrix.ReadNum()
	fmt.Println("Введите матрицу: ")
	amat := matrix.ReadMatrix(an, am)

	fmt.Println("\n Матрица B")
	fmt.Print("Введите количество строк: ")
	bn := matrix.ReadNum()

	if bn != am {
		fmt.Println("Количество строк матриц не совпадает")
		return
	}
	fmt.Print("Введите количество столбцов: ")
	bm := matrix.ReadNum()
	fmt.Println("Введите матрицу: ")
	bmat := matrix.ReadMatrix(bn, bm)

	fmt.Println("\n Резулитирующая матрица")
	fmt.Println("Обычное умножение: ")
	smmat := matrix.SimpleMult(amat, bmat)
	smmat.PrintMatrix()

	fmt.Println("\n Умножение по Винограду")
	wmmat := matrix.WinogradMult(amat, bmat)
	wmmat.PrintMatrix()

	fmt.Println("Умножение по Винограду (улучшенное): ")
	wimmat := matrix.WinogradMult(amat, bmat)
	wimmat.PrintMatrix()
}