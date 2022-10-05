package matrix

import (
	"testing"
)

// Simple multiplication benchmarks.

func BenchmarkSimple100(b *testing.B) {
	N := 100
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple200(b *testing.B) {
	N := 200
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple300(b *testing.B) {
	N := 300
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple400(b *testing.B) {
	N := 400
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple500(b *testing.B) {
	N := 500
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple600(b *testing.B) {
	N := 600
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple700(b *testing.B) {
	N := 700
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple800(b *testing.B) {
	N := 800
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

// Winograd multiplication benchmarks.

func BenchmarkWinograd100(b *testing.B) {
	N := 100
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd200(b *testing.B) {
	N := 200
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd300(b *testing.B) {
	N := 300
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd400(b *testing.B) {
	N := 400
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd500(b *testing.B) {
	N := 500
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd600(b *testing.B) {
	N := 600
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd700(b *testing.B) {
	N := 700
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd800(b *testing.B) {
	N := 800
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

// Winograd improved multiplication benchmarks.

func BenchmarkWinogradImp100(b *testing.B) {
	N := 100
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp200(b *testing.B) {
	N := 200
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp300(b *testing.B) {
	N := 300
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp400(b *testing.B) {
	N := 400
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp500(b *testing.B) {
	N := 500
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp600(b *testing.B) {
	N := 600
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp700(b *testing.B) {
	N := 700
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp800(b *testing.B) {
	N := 800
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkSimple101(b *testing.B) {
	N := 101
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple201(b *testing.B) {
	N := 201
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple301(b *testing.B) {
	N := 301
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple401(b *testing.B) {
	N := 401
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple501(b *testing.B) {
	N := 501
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple601(b *testing.B) {
	N := 601
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple701(b *testing.B) {
	N := 701
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

func BenchmarkSimple801(b *testing.B) {
	N := 801
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		SimpleMult(amat, bmat)
	}
}

// Winograd multiplication benchmarks.

func BenchmarkWinograd101(b *testing.B) {
	N := 101
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd201(b *testing.B) {
	N := 201
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd301(b *testing.B) {
	N := 301
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd401(b *testing.B) {
	N := 401
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd501(b *testing.B) {
	N := 501
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd601(b *testing.B) {
	N := 601
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd701(b *testing.B) {
	N := 701
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

func BenchmarkWinograd801(b *testing.B) {
	N := 801
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMult(amat, bmat)
	}
}

// Winograd improved multiplication benchmarks.

func BenchmarkWinogradImp101(b *testing.B) {
	N := 101
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp201(b *testing.B) {
	N := 201
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp301(b *testing.B) {
	N := 301
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp401(b *testing.B) {
	N := 401
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp501(b *testing.B) {
	N := 501
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp601(b *testing.B) {
	N := 601
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp701(b *testing.B) {
	N := 701
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkWinogradImp801(b *testing.B) {
	N := 801
	amat := formResMat(N, N)
	randomFill(amat, 101)
	bmat := formResMat(N, N)
	randomFill(bmat, 101)

	for i := 0; i < b.N; i++ {
		WinogradMultImp(amat, bmat)
	}
}

func BenchmarkRowsCols100(b *testing.B) {
	N := 100
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		precomputeWinogradRows(amat)
		precomputeWinogradCols(bmat)
	}
}

func BenchmarkRowsCols200(b *testing.B) {
	N := 200
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		precomputeWinogradRows(amat)
		precomputeWinogradCols(bmat)
	}
}
func BenchmarkRowsCols300(b *testing.B) {
	N := 300
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		precomputeWinogradRows(amat)
		precomputeWinogradCols(bmat)
	}
}
func BenchmarkRowsCols400(b *testing.B) {
	N := 400
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		precomputeWinogradRows(amat)
		precomputeWinogradCols(bmat)
	}
}
func BenchmarkRowsCols500(b *testing.B) {
	N := 500
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		precomputeWinogradRows(amat)
		precomputeWinogradCols(bmat)
	}
}
func BenchmarkRowsCols600(b *testing.B) {
	N := 600
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		precomputeWinogradRows(amat)
		precomputeWinogradCols(bmat)
	}
}

func BenchmarkRowsCols700(b *testing.B) {
	N := 700
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		precomputeWinogradRows(amat)
		precomputeWinogradCols(bmat)
	}
}

func BenchmarkRowsCols800(b *testing.B) {
	N := 800
	amat := formResMat(N, N)
	randomFill(amat, 100)
	bmat := formResMat(N, N)
	randomFill(bmat, 100)

	for i := 0; i < b.N; i++ {
		precomputeWinogradRows(amat)
		precomputeWinogradCols(bmat)
	}
}
