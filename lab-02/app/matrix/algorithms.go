package matrix

// SimpleMult used to multiply two matrixes using single multiplication method.
func SimpleMult(amat, bmat MInt) MInt {
	rmat := formResMat(amat.n, bmat.m)

	for i := 0; i < rmat.n; i++ {
		for j := 0; j < rmat.m; j++ {
			for k := 0; k < amat.m; k++ {
				rmat.mat[i][j] += amat.mat[i][k] * bmat.mat[k][j]
			}
		}
	}

	return rmat
}

// WinogradMult used to multiply two matrixes using Winograd method.
func WinogradMult(amat, bmat MInt) MInt {
	rmat := formResMat(amat.n, bmat.m)

	rowcf := precomputeWinogradRows(amat)
	colcf := precomputeWinogradCols(bmat)

	for i := 0; i < rmat.n; i++ {
		for j := 0; j < rmat.m; j++ {
			for k := 0; k < rmat.m/2; k++ {
				rmat.mat[i][j] = rmat.mat[i][j] + (amat.mat[i][k*2]+bmat.mat[k*2+1][j])*(amat.mat[i][k*2+1]+bmat.mat[k*2][j]) -
					rowcf.mat[i][k] - colcf.mat[k][j]
			}
		}
	}

	if rmat.m%2 != 0 {
		for i := 0; i < rmat.n; i++ {
			for j := 0; j < rmat.m; j++ {
				rmat.mat[i][j] += amat.mat[i][amat.m-1] * bmat.mat[bmat.n-1][j]
			}
		}
	}

	return rmat
}

func precomputeWinogradRows(mat MInt) MInt {
	s := mat.m / 2
	cf := formResMat(mat.n, s)

	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m/2; j++ {
			cf.mat[i][j] = mat.mat[i][j*2] * mat.mat[i][j*2+1]
		}
	}

	return cf
}

func precomputeWinogradCols(mat MInt) MInt {
	s := mat.n / 2
	cf := formResMat(s, mat.m)

	for i := 0; i < mat.n/2; i++ {
		for j := 0; j < mat.m; j++ {
			cf.mat[i][j] = mat.mat[i*2][j] * mat.mat[i*2+1][j]
		}
	}

	return cf
}

// WinogradMultImp used to multiply two matrixes using impproved Winograd method.
func WinogradMultImp(amat, bmat MInt) MInt {
	rmat := formResMat(amat.n, bmat.m)

	rowcf := precomputeWinogradRowsImp(amat)
	colcf := precomputeWinogradColsImp(bmat)

	for i := 0; i < rmat.n; i++ {
		for j := 0; j < rmat.m; j++ {
			for k := 0; k < rmat.m-1; k += 2 {
				l := k/2 + k%2
				rmat.mat[i][j] += (amat.mat[i][k]+bmat.mat[k+1][j])*(amat.mat[i][k+1]+bmat.mat[k][j]) -
					rowcf.mat[i][l] - colcf.mat[l][j]
			}
		}
	}

	if rmat.m%2 != 0 {
		k := amat.m - 1
		for i := 0; i < rmat.n; i++ {
			for j := 0; j < rmat.m; j++ {
				rmat.mat[i][j] += amat.mat[i][k] * bmat.mat[k][j]
			}
		}
	}

	return rmat
}

func precomputeWinogradRowsImp(mat MInt) MInt {
	s := mat.m / 2
	cf := formResMat(mat.n, s)

	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m-1; j += 2 {
			cf.mat[i][j/s+j%2] = mat.mat[i][j] * mat.mat[i][j+1]
		}
	}

	return cf
}

func precomputeWinogradColsImp(mat MInt) MInt {
	s := mat.n / 2
	cf := formResMat(s, mat.m)

	for i := 0; i < mat.n-1; i += 2 {
		for j := 0; j < mat.m; j++ {
			cf.mat[i/s+i%2][j] = mat.mat[i][j] * mat.mat[i+1][j]
		}
	}

	return cf
}
