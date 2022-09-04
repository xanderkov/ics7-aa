package levenstein

// import "math"

func CountLevNoRec(src, dest string) (int, MInt) {
	
	srcRune, destRune := []rune(src), []rune(dest)
	n, m := len(srcRune), len(destRune)

	distMat := makeMatrix(n, m)

	for i := 1; i < n + 1; i++ {
		for j := 1; j < m + 1; j++ {
			insDist := distMat[i][j - 1] + 1
			delDist := distMat[i - 1][j] + 1
			
			match := 1 // Штраф за Match, если надо менять
			
			if srcRune[i - 1] == destRune[j - 1] {
				match = 0
			}
			eqDist := distMat[i - 1][j - 1] + match

			dist := getMinOfValues(insDist, delDist, eqDist)
			distMat[i][j] = dist
		}
	}
	shortDist := distMat[n][m]
	return shortDist, distMat
}

func CountDamNoRec(src, dest string) (int, MInt) {
	var (n, m, dist, shortDist, transDist int)

	srcRune, destRune := []rune(src), []rune(dest)

	n, m = len(srcRune), len(destRune)

	distMat := makeMatrix(n, m)

	for i := 1; i < n + 1; i++ {
		for j := 1; j < m + 1; j++ {
			insDist := distMat[i][j - 1] + 1
			delDist := distMat[i - 1][j] + 1
			
			match := 1
			if src[i - 1] == dest[j - 1] {
				match = 0
			}
			eqDist := distMat[i - 1][j - 1] + match

			transDist = -1
			if i > 1 && j > 1 {
				transDist = distMat[i - 2][j - 2] + 1
			}

			if transDist != -1 && srcRune[i - 1] == destRune[j - 2] && srcRune[i - 2] == destRune[j - 1] {
				dist = getMinOfValues(insDist, delDist, eqDist, transDist)
			} else {
				dist = getMinOfValues(insDist, delDist, eqDist)
			}
			distMat[i][j] = dist
		}
	}

	shortDist = distMat[n][m]

	return shortDist, distMat
}

func _countDamRecElem(src, dest []rune, i, j int) int {
	if (getMinOfValues(i, j) == 0) {
		return getMaxOf2Values(i, j)
	}
	
	match := 1
	if (src[i - 1] == dest[j - 1]) {
		match = 0
	}
 
	insert := _countDamRecElem(src, dest, i, j - 1) + 1
	delete := _countDamRecElem(src, dest, i - 1, j) + 1
	replace := match + _countDamRecElem(src, dest, i - 1, j - 1)

	transpose := -1

	if i > 1 && j > 1 {
		transpose = _countDamRecElem(src, dest, i - 2, j - 2) + 1
	}

	min := 0
	if transpose != -1 && src[i - 1] == dest[j - 2] && src[i - 2] == dest[j - 1] {
		min = getMinOfValues(insert, delete, replace, transpose)
	} else {
		min = getMinOfValues(insert, delete, replace)
	}
	return min
}

func CountDamRecNoCache(src, dest string) int {

	srcRune, destRune := []rune(src), []rune(dest)
	n, m := len(srcRune), len(destRune)

	return _countDamRecElem(srcRune, destRune, n, m)
}

func _countDamRecElemCache(src, dest []rune, i, j int, distMat MInt) int {
	if (getMinOfValues(i, j) == 0) {
		return getMaxOf2Values(i, j)
	}

	if distMat[i][j] != -1 {
		return distMat[i][j]
	}
	
	match := 1
	if (src[i - 1] == dest[j - 1]) {
		match = 0
	}
 
	insert := _countDamRecElemCache(src, dest, i, j - 1, distMat) + 1
	delete := _countDamRecElemCache(src, dest, i - 1, j, distMat) + 1
	replace := match + _countDamRecElemCache(src, dest, i - 1, j - 1, distMat)

	transpose := -1

	if i > 1 && j > 1 {
		transpose = _countDamRecElemCache(src, dest, i - 2, j - 2, distMat) + 1
	}

	min := 0
	if transpose != -1 && src[i - 1] == dest[j - 2] && src[i - 2] == dest[j - 1] {
		min = getMinOfValues(insert, delete, replace, transpose)
	} else {
		min = getMinOfValues(insert, delete, replace)
	}

	distMat[i][j] = min
	return distMat[i][j]

}

func CountDamRecCache(src, dest string) int {
	srcRune, destRune := []rune(src), []rune(dest)
	n, m := len(srcRune), len(destRune)

	distMat := makeMatrixInf(n, m)

	return _countDamRecElemCache(srcRune, destRune, n, m, distMat)
}