package levenstein

// import "math"

func CountLevNoRec(src, dest string) int {
	return len(src) + len(dest)
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
			eq := 1
			if src[i - 1] == dest[j - 1] {
				eq = 0
			}
			eqDist := distMat[i - 1][j - 1] + eq
			transDist = -1
			if i > 1 && j > 1 {
				transDist = distMat[i - 2][j - 2] + 1
			}

			if transDist != -1 && srcRune[i - 1] == destRune[j - 2] && srcRune[i - 2] == destRune[j - 1] {
				dist = getMinOfMins(insDist, delDist, eqDist, transDist)
			} else {
				dist = getMinOfMins(insDist, delDist, eqDist)
			}
			distMat[i][j] = dist
		}
	}

	shortDist = distMat[n][m]

	return shortDist, distMat
}

func CountDamRecNoCache(src, dest string) int {
	return len(src) + len(dest)
}

func CountDamRecCache(src, dest string) int {
	return len(src) + len(dest)
}