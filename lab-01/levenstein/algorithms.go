package levenstein

// import "math"

func CountLevNoRec(src, dest string) int {
	return len(src) + len(dest)
}

func CountDamNoRec(src, dest string) int {
	var (n, m, dist, shortDist, transDist int)

	src, dest := []rune(src), []rune(dest)

	n, m = len(src), len(dest)

	distMat := makeMatrix(n, m)

	for i := 1; i < n + 1; i++ {
		for j := 1
	}

	return len(src) + len(dest)
}

func CountDamRecNoCache(src, dest string) int {
	return len(src) + len(dest)
}

func CountDamRecCache(src, dest string) int {
	return len(src) + len(dest)
}