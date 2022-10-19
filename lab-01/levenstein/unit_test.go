package levenstein

import "testing"

func TestEmptyString(t *testing.T) {
	eth := 0
	w1 := ""
	w2 := ""
	
	res, mat := CountLevNoRec(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res, mat = CountDamNoRec(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res = CountDamRecNoCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = CountDamRecCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

}

func TestString1(t *testing.T) {
	eth := 1
	w1 := "book"
	w2 := "bosk"
	
	res, mat := CountLevNoRec(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res, mat = CountDamNoRec(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res = CountDamRecNoCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = CountDamRecCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

}

func TestString2(t *testing.T) {
	eth := 2
	w1 := "book"
	w2 := "back"
	
	res, mat := CountLevNoRec(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res, mat = CountDamNoRec(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res = CountDamRecNoCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = CountDamRecCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

}

func TestString3(t *testing.T) {
	eth := 3
	w1 := "book"
	w2 := "bacc"
	
	res, mat := CountLevNoRec(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res, mat = CountDamNoRec(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res = CountDamRecNoCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = CountDamRecCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

}

func TestString4(t *testing.T) {
	eth := 4
	w1 := "aboba"
	w2 := "acacb"
	
	res, mat := CountLevNoRec(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res, mat = CountDamNoRec(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res = CountDamRecNoCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = CountDamRecCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

}

func TestString21(t *testing.T) {
	eth := 2
	w1 := "дверь"
	w2 := "деврь"
	
	res, mat := CountLevNoRec(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	eth = 1

	res, mat = CountDamNoRec(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res = CountDamRecNoCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = CountDamRecCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

}

func TestString0(t *testing.T) {
	eth := 0
	w1 := "дверь"
	w2 := "дверь"
	
	res, mat := CountLevNoRec(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res, mat = CountDamNoRec(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res = CountDamRecNoCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = CountDamRecCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

}