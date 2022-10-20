package main

/*
#include <pthread.h>
#include <time.h>
#include <stdio.h>

static long long getThreadCpuTimeNs() {
    struct timespec t;
    if (clock_gettime(CLOCK_THREAD_CPUTIME_ID, &t)) {
        perror("clock_gettime");
        return 0;
    }
    return t.tv_sec * 1000000000LL + t.tv_nsec;
}
*/
import "C"

import (
	"fmt"
	"./matrix"
)

func test(amat, bmat matrix.MInt, n int){
	cputime1 := C.getThreadCpuTimeNs()
	
	for i := 0; i < n; i++ {
		matrix.WinogradMult(amat, bmat)
	}
	
	cputime2 := C.getThreadCpuTimeNs()
	fmt.Printf("CPU time = %d ns\n", (cputime2 - cputime1))
	
}

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
	test(amat, bmat, 10)
}